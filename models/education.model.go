package models

import (
	"time"

	"gorm.io/gorm"
)

type Education struct {
	gorm.Model
	Institution  string    `json:"institution" gorm:"not null"`
	Degree       string    `json:"degree" gorm:"not null"`
	FieldOfStudy string    `json:"field_of_study" gorm:"not null"`
	StartDate    time.Time `json:"start_date" gorm:"not null"`
	EndDate      time.Time `json:"end_date"`
	Skills       string    `json:"skills" gorm:"not null"`
	UserID       uint      `json:"user_id" gorm:"not null"`
}
