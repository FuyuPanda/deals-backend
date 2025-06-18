package models

import "time"

type employeeReimbursement struct {
	Id         uint `gorm:"primaryKey"`
	employeeID int
	createdBy  int
	updatedBy  int
	Points     int64

	Amount      int    `json:"amount"`
	Description string `json:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
