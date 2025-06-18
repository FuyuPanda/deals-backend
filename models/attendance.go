package models

import "Time"

type attendance struct {
	id        uint `gorm:"primaryKey"`
	createdBy int
	updatedBy int
	Points    int64

	checkInTime  time.Time
	checkOutTime time.Time
	totalHours   int
	createdOn    time.Time
	updatedOn    time.Time
}
