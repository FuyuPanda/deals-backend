package models

import "time"

type employeeActivities struct {
	Id         uint `gorm:"primaryKey"`
	EmployeeID int
	Points     int64

	Ip          string
	Description string
	Device      string

	CreatedAt time.Time
}
