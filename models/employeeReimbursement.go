package models

import "time"

type EmployeeReimbursement struct {
	Id         uint `gorm:"primaryKey"`
	EmployeeID int
	CreatedBy  int
	UpdatedBy  int

	Amount              int       `json:"amount"`
	Description         string    `json:"description"`
	TimeOfReimbursement time.Time `json:"time_of_reimbursement"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
