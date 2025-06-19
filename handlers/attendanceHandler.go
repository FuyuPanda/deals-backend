package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/FuyuPanda/deals-backend/db"
	"github.com/FuyuPanda/deals-backend/middleware"
	"github.com/FuyuPanda/deals-backend/models"
)

func CheckInAttendance(w http.ResponseWriter, r *http.Request) {
	var attendances models.Attendance
	if err := json.NewDecoder(r.Body).Decode(&attendances); err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	userID, ok := r.Context().Value(middleware.UserIDKey).(uint)
	if !ok {
		http.Error(w, "Invalid user ID", http.StatusInternalServerError)
		return
	}

	now := time.Now()
	weekday := now.Weekday()

	if weekday == time.Saturday || weekday == time.Sunday {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Cannot clock in during weekends",
			"status":  "Failed",
		})
		return

	}

	attendances.EmployeeID = int(userID)
	attendances.CheckInTime = time.Now()
	attendances.CreatedBy = int(userID)
	attendances.UpdatedBy = int(userID)

	if err := db.DB.Create(&attendances).Error; err != nil {
		http.Error(w, "Failed to Clock In", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(attendances)

}

func CheckOutAttendance(w http.ResponseWriter, r *http.Request) {
	var attendances models.Attendance
	if err := json.NewDecoder(r.Body).Decode(&attendances); err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	userID, ok := r.Context().Value(middleware.UserIDKey).(uint)
	if !ok {
		http.Error(w, "Invalid user ID", http.StatusInternalServerError)
		return
	}

	if err := db.DB.Where("employee_id = ? AND DATE(check_in_time) = CURRENT_DATE", userID).
		First(&attendances).Error; err != nil {
		http.Error(w, "Attendance record not found", http.StatusNotFound)
		return
	}

	totalHours := time.Since(attendances.CheckInTime).Round(time.Hour).Hours()

	// Update checkout and total hours
	if err := db.DB.Model(&attendances).Updates(map[string]interface{}{
		"CheckOutTime": time.Now(),
		"TotalHours":   totalHours,
		"UpdatedBy":    userID,
	}).Error; err != nil {
		http.Error(w, "Failed to Clock Out", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(attendances)

}

func GetAttendanceHistory(w http.ResponseWriter, r *http.Request) {
	var attendances []models.Attendance
	userID, ok := r.Context().Value(middleware.UserIDKey).(uint)
	if !ok {
		http.Error(w, "Invalid user ID", http.StatusInternalServerError)
		return
	}

	if err := db.DB.Where("employee_id = ?", userID).Find(&attendances).Error; err != nil {
		http.Error(w, "Attendance record not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(attendances)

}
