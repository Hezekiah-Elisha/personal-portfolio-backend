package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Username  string         `json:"username" gorm:"unique;not null"`
	Email     string         `json:"email" gorm:"unique;not null"`
	Password  string         `json:"password" gorm:"not null"`
	Role      string         `json:"role" gorm:"default:'user'"`
	CreatedAt time.Time      `json:"created_at" gorm:"not null autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"not null autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index null"`
}
