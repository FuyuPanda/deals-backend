package models

import "time"

type Attendance struct {
	Id         uint `gorm:"primaryKey"`
	EmployeeID int
	CreatedBy  int
	UpdatedBy  int

	CheckInTime  time.Time
	CheckOutTime time.Time
	TotalHours   int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type AttendanceResponse struct {
	EmployeeName string `json:"employee_name"`

	CheckInTime  time.Time `json:"check_in_time"`
	CheckOutTime time.Time `json:"check_out_time"`
	TotalHours   int       `json:"total_hours"`
	Amount       int       `json:"amount"`
}
