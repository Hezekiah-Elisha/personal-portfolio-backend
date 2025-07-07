package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Title       string  `json:"title" gorm:"not null"`
	Description string  `json:"description" gorm:"not null"`
	TechStack   string  `json:"tech_stack" gorm:"not null"`
	Link        *string `json:"link"`
	SourceCode  *string `json:"source_code"`
	Image       string  `json:"image" gorm:"not null"`
	UserID      uint    `json:"user_id" gorm:"not null"`
}
