package controllers

import "github.com/gin-gonic/gin"

type CreateExperienceInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Company     string `json:"company" binding:"required"`
	Location    string `json:"location" binding:"required"`
	StartDate   string `json:"start_date" binding:"required"`
	EndDate     string `json:"end_date"`
	UserID      uint   `json:"user_id" binding:"required"`
}

func CreateExperience(c *gin.Context) {

}
