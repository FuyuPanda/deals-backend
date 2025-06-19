package models

import "time"

type PayslipReimbursementDetails struct {
	Id              uint `gorm:"primaryKey"`
	PayslipID       int
	ReimbursementID int
	CreatedBy       int
	UpdatedBy       int

	CreatedAt time.Time
	UpdatedAt time.Time
}
