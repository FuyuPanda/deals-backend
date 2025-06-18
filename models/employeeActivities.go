package models

import "Time"

type employeeActivities struct {
	id         uint `gorm:"primaryKey"`
	employeeID int
	Points     int64

	ip          string
	description string
	device      string
}
