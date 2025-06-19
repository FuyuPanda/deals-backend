package models

import "time"

type PayslipOvertimeDetails struct {
	Id         uint `gorm:"primaryKey"`
	PayslipID  int
	OvertimeID int
	CreatedBy  int
	UpdatedBy  int

	CreatedAt time.Time
	UpdatedAt time.Time
}
