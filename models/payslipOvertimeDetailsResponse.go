package models

import "time"

type PayslipOvertimeDetailsResponse struct {
	EmployeeName string    `json:"employee_name"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
	TotalHours   int       `json:"total_hours"`
	Amount       int       `json:"amount"`
}
