package models

import "Time"

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
	createdOn    time.Time
	updatedOn    time.Time
}
