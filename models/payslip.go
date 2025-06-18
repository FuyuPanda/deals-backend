package models

import "time"

type payslip struct {
	id         uint `gorm:"primaryKey"`
	employeeID int
	createdBy  int
	updatedBy  int
	Points     int64

	isDownloaded bool
	status       string `json:"status"`
	downloadedOn time.Time
	period       time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
