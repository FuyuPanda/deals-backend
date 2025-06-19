package models

import "time"

type PayslipReimbursementDetailsResponse struct {
	EmployeeName        string    `json:"employee_name"`
	Amount              int       `json:"amount"`
	Description         string    `json:"description"`
	TimeOfReimbursement time.Time `json:"time_of_reimbursement"`
}
