package models

import "gorm.io/gorm"

type Education struct {
	gorm.Model
	Institution  string `json:"institution" gorm:"not null"`
	Degree       string `json:"degree" gorm:"not null"`
	FieldOfStudy string `json:"field_of_study" gorm:"not null"`
	StartDate    string `json:"start_date" gorm:"not null"`
	EndDate      string `json:"end_date"`
	Skills       string `json:"skills" gorm:"not null"`
	UserID       uint   `json:"user_id" gorm:"not null"`
}
