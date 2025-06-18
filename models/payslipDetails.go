package models

import "time"

type payslipDetails struct {
	id              uint `gorm:"primaryKey"`
	overtimeID      int
	reimbursementID int
	createdBy       int
	updatedBy       int
	Points          int64

	CreatedAt time.Time
	UpdatedAt time.Time
}
