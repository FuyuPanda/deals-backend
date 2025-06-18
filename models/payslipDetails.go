package models

import "Time"

type payslipDetails struct {
	id              uint `gorm:"primaryKey"`
	overtimeID      int
	reimbursementID int
	createdBy       int
	updatedBy       int
	Points          int64

	createdOn time.Time
	updatedOn time.Time
}
