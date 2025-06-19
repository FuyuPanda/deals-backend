package models

import "time"

type Employees struct {
	Id        uint `gorm:"primaryKey"`
	CreatedBy int
	UpdatedBy int

	Username string  `json:"username"`
	Password string  `json:"password"`
	FullName string  `json:"full_name"`
	Salary   float64 `json:"salary"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
