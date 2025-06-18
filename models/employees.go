package models

import "Time"

type Employees struct {
	id        uint `json:"id" gorm:"primaryKey"`
	createdBy int
	updatedBy int
	Points    int64

	username string `json:"username"`
	password string `json:"password"`
	fullName string `json:"full_name"`
	salary   string `json:"salary"`

	createdOn time.Time
	updatedOn time.Time
}
