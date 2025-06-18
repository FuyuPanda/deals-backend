package models

import "Time"

type employeeReimbursement struct {
	id         uint `gorm:"primaryKey"`
	employeeID int
	createdBy  int
	updatedBy  int
	Points     int64

	amount      int    `json:"amount"`
	description string `json:"description"`
	createdOn   time.Time
	updatedOn   time.Time
}
