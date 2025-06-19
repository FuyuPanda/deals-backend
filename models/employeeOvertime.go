package models

import "time"

type EmployeeOvertime struct {
	Id         uint `gorm:"primaryKey"`
	CreatedBy  int
	UpdatedBy  int
	EmployeeID int

	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	TotalHours int
}
