package models

import (
	"time"
)

type Payslip struct {
	Id         uint `gorm:"primaryKey"`
	EmployeeID int  `json:"employee_id"`
	CreatedBy  int
	UpdatedBy  int

	IsDownloaded bool
	Status       string `json:"status"`
	TotalAmount  int
	DownloadedOn time.Time
	Period       time.Time `json:"period"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
