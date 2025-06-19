package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/FuyuPanda/deals-backend/db"
	"github.com/FuyuPanda/deals-backend/middleware"
	"github.com/FuyuPanda/deals-backend/models"
)

func CreatePayroll(w http.ResponseWriter, r *http.Request) {
	username, ok := r.Context().Value(middleware.UsernameKey).(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid user ID",
			"status":  "Failed",
		})
		return
	}

	if username != "admin" {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Forbidden Access",
			"status":  "Failed",
		})
		return
	}

	userID, ok := r.Context().Value(middleware.UserIDKey).(uint)
	if !ok {
		http.Error(w, "Invalid user ID", http.StatusInternalServerError)
		return
	}

	var payslips models.Payslip
	if err := json.NewDecoder(r.Body).Decode(&payslips); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid Input",
			"status":  "Failed",
		})
		return
	}
	start := middleware.FirstOfMonth(payslips.Period)
	end := middleware.EndOfMonth(payslips.Period)
	fmt.Println(payslips.Period)

	if err := db.DB.Where("employee_id = ? AND DATE(period) BETWEEN ? AND ?", payslips.EmployeeID, start, end).First(&payslips).Error; err == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Payroll for this period has been processed!",
			"status":  "Failed",
		})
		return
	}

	amount := 0

	//get salary of that employee
	var employee models.Employees

	if err := db.DB.Where("id = ?", payslips.EmployeeID).First(&employee).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Employee Data Not Found",
			"status":  "Failed",
		})
		return
	}
	salary := int(employee.Salary)
	salaryPerHour := salary / 8

	//get List of Attendances
	var attendances []models.Attendance
	if err := db.DB.Where("employee_id = ? AND DATE(check_in_time) >= ? AND DATE(check_out_time) <= ?", payslips.EmployeeID, start, end).
		Find(&attendances).Error; err == nil {
		for _, attendance := range attendances {
			//get salary per hour

			amount += salaryPerHour * attendance.TotalHours
		}

	}
	payslips.TotalAmount = amount
	payslips.IsDownloaded = false
	payslips.Period = end
	payslips.CreatedBy = int(userID)
	payslips.UpdatedBy = int(userID)

	if err := db.DB.Create(&payslips).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Failed to process payroll",
			"status":  "Failed",
		})
		return
	}

	//get List of Overtimes
	var overtimes []models.EmployeeOvertime
	if err := db.DB.Where("employee_id = ? AND DATE(start_time) >= ? AND DATE(end_time) <= ?", payslips.EmployeeID, start, end).
		Find(&overtimes).Error; err == nil {
		for _, overtime := range overtimes {
			var overtimePayslip models.PayslipOvertimeDetails

			overtimePayslip.OvertimeID = int(overtime.Id)
			overtimePayslip.PayslipID = int(payslips.Id)
			overtimePayslip.CreatedBy = int(userID)
			overtimePayslip.UpdatedBy = int(userID)
			amount += salaryPerHour * overtime.TotalHours

			if err := db.DB.Create(&overtimePayslip).Error; err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{
					"message": "Failed to process payroll",
					"status":  "Failed",
				})
				return
			}

		}
	}

	//get List of Reimbursement
	var reimbursements []models.EmployeeReimbursement
	if err := db.DB.Where("employee_id = ? AND DATE(time_of_reimbursement) >= ? AND DATE(time_of_reimbursement)<= ?", payslips.EmployeeID, start, end).
		Find(&reimbursements).Error; err == nil {
		for _, reimbursement := range reimbursements {
			var reimbursementPayslip models.PayslipReimbursementDetails

			reimbursementPayslip.ReimbursementID = int(reimbursement.Id)
			reimbursementPayslip.PayslipID = int(payslips.Id)
			reimbursementPayslip.CreatedBy = int(userID)
			reimbursementPayslip.UpdatedBy = int(userID)

			amount += reimbursement.Amount

			if err := db.DB.Create(&reimbursementPayslip).Error; err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{
					"message": "Failed to process payroll",
					"status":  "Failed",
				})
				return
			}
		}
	}

	//update total amount in payroll
	amount += payslips.TotalAmount
	err := db.DB.Model(&payslips).Update("TotalAmount", amount).Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Failed to process payroll",
			"status":  "Failed",
		})
		return
	}

	json.NewEncoder(w).Encode(payslips)

}

func GetPayroll(w http.ResponseWriter, r *http.Request) {
	username, ok := r.Context().Value(middleware.UsernameKey).(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid user ID",
			"status":  "Failed",
		})
		return
	}

	if username == "admin" {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Forbidden Access",
			"status":  "Failed",
		})
		return
	}

	userID, ok := r.Context().Value(middleware.UserIDKey).(uint)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid User ID",
			"status":  "Failed",
		})
		return
	}

	//get salary of that employee
	var employee models.Employees

	if err := db.DB.Where("id = ?", int(userID)).First(&employee).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Employee Data Not Found",
			"status":  "Failed",
		})
		return
	}
	salary := int(employee.Salary)
	salaryPerHour := salary / 8

	var payslips []models.Payslip
	var responses []models.PayslipResponse
	if err := db.DB.Where("employee_id = ?", int(userID)).
		Find(&payslips).Error; err == nil {
		for _, payslip := range payslips {
			// Get overtime details
			var overtimes []models.PayslipOvertimeDetailsResponse
			err := db.DB.Table("payslips").
				Joins("JOIN payslip_overtime_details ON payslip_overtime_details.payslip_id = payslips.id").
				Joins("JOIN employee_overtimes on employee_overtimes.id = payslip_overtime_details.overtime_id").
				Joins("JOIN employees on employee_overtimes.employee_id = employees.id").
				Where("employees.id = ?", int(userID)).
				Select("employees.full_name as employee_name, employee_overtimes.start_time, employee_overtimes.end_time, employee_overtimes.total_hours, 0 as amount").
				Scan(&overtimes).Error

			if err == nil {
				for i := range overtimes {
					overtimes[i].Amount = salaryPerHour * overtimes[i].TotalHours

				}

			}
			//  Get reimbursement details
			var reimbursements []models.PayslipReimbursementDetailsResponse
			err = db.DB.Table("payslips").
				Joins("JOIN payslip_reimbursement_details ON payslip_reimbursement_details.payslip_id = payslips.id").
				Joins("JOIN employee_reimbursements on employee_reimbursements.id = payslip_reimbursement_details.reimbursement_id").
				Joins("JOIN employees on employee_reimbursements.employee_id = employees.id").
				Where("employees.id = ?", int(userID)).
				Select("employees.full_name as employee_name, employee_reimbursements.amount,employee_reimbursements.description,employee_reimbursements.time_of_reimbursement").
				Scan(&reimbursements).Error
			if err != nil {

			}

			//Get Attendances
			var attendances []models.AttendanceResponse
			start := middleware.FirstOfMonth(payslip.Period)
			end := middleware.EndOfMonth(payslip.Period)
			err = db.DB.Table("payslips").
				Joins("JOIN attendances ON attendances.employee_id = payslips.employee_id").
				Joins("JOIN employees on attendances.employee_id = employees.id").
				Where("employees.id = ? AND DATE(attendances.check_in_time) >= ? AND DATE(attendances.check_out_time) <= ?", int(userID), start, end).
				Select("employees.full_name as employee_name, attendances.check_in_time, attendances.check_out_time, attendances.total_hours, 0 as amount").
				Scan(&attendances).Error
			if err == nil {
				for i := range attendances {
					attendances[i].Amount = salaryPerHour * attendances[i].TotalHours

				}

			}

			// Build response object
			response := models.PayslipResponse{
				ID:             payslip.Id,
				EmployeeID:     payslip.EmployeeID,
				CreatedBy:      payslip.CreatedBy,
				UpdatedBy:      payslip.UpdatedBy,
				IsDownloaded:   payslip.IsDownloaded,
				Status:         payslip.Status,
				TotalAmount:    payslip.TotalAmount,
				DownloadedOn:   payslip.DownloadedOn,
				Period:         payslip.Period,
				CreatedAt:      payslip.CreatedAt,
				UpdatedAt:      payslip.UpdatedAt,
				Overtimes:      overtimes,
				Reimbursements: reimbursements,
				Attendances:    attendances,
			}

			responses = append(responses, response)
		}

	}
	json.NewEncoder(w).Encode(responses)

}

func GetSummaryPayrollEmployee(w http.ResponseWriter, r *http.Request) {
	var payslips []models.SummaryPayrollEmployeeResponse
	username, ok := r.Context().Value(middleware.UsernameKey).(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid user ID",
			"status":  "Failed",
		})
		return
	}

	if username != "admin" {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Forbidden Access",
			"status":  "Failed",
		})
		return
	}
	err := db.DB.Table("payslips").
		Joins("JOIN employees on payslips.employee_id = employees.id").
		Select("employees.username as employee_name, SUM(payslips.total_amount) as amount").
		Group("employees.username").
		Scan(&payslips).Error
	if err == nil {
		for i := range payslips {
			var employee models.Employees

			if err := db.DB.Where("username = ?", payslips[i].EmployeeName).First(&employee).Error; err == nil {
				payslips[i].EmployeeName = employee.FullName
			}

		}

	}

	json.NewEncoder(w).Encode(payslips)
}

func GetTotalPayrollEmployee(w http.ResponseWriter, r *http.Request) {
	var payslips models.TotalPayrollResponse
	username, ok := r.Context().Value(middleware.UsernameKey).(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid user ID",
			"status":  "Failed",
		})
		return
	}

	if username != "admin" {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Forbidden Access",
			"status":  "Failed",
		})
		return
	}
	err := db.DB.Table("payslips").
		Select("SUM(payslips.total_amount) as amount").
		Scan(&payslips).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Data not found",
			"status":  "Failed",
		})
		return

	}

	json.NewEncoder(w).Encode(payslips)
}
