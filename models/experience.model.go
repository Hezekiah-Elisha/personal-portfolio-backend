package models

type Experience struct {
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	Company     string `json:"company" gorm:"not null"`
	Location    string `json:"location" gorm:"not null"`
	StartDate   string `json:"start_date" gorm:"not null"`
	EndDate     string `json:"end_date"`
	UserID      uint   `json:"user_id" gorm:"not null"`
}
