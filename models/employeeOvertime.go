package models

import "time"

type employeeOvertime struct {
	Id         uint `gorm:"primaryKey"`
	createdBy  int
	updatedBy  int
	employeeID int
	Points     int64

	startTime  time.Time
	endTime    time.Time
	totalHours int
}
