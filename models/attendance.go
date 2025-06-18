package models

import "time"

type attendance struct {
	id        uint `gorm:"primarykey"`
	CreatedBy int
	UpdatedBy int
	Points    int64

	CheckInTime  time.Time
	CheckOutTime time.Time
	TotalHours   int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
