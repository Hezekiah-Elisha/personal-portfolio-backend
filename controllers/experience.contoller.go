package controllers

import (
	"log"
	"net/http"
	"personal-portfolio-backend/config"
	"personal-portfolio-backend/models"

	"github.com/gin-gonic/gin"
)

type CreateExperienceInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Company     string `json:"company" binding:"required"`
	Location    string `json:"location" binding:"required"`
	StartDate   string `json:"start_date" binding:"required"`
	EndDate     string `json:"end_date"`
	UserID      uint   `json:"user_id"`
}

type EditExperienceInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Company     string `json:"company"`
	Location    string `json:"location"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

type ExperienceResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Company     string `json:"company"`
	Location    string `json:"location"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	UserID      uint   `json:"user_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func CreateExperience(c *gin.Context) {
	var input CreateExperienceInput

	if err := c.ShouldBindJSON(&input); err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data", "details": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	// Convert to uint (assuming your user ID is uint)
	log.Println("User ID from context:", userID)
	userIDUint, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	experience := models.Experience{
		Title:       input.Title,
		Description: input.Description,
		Company:     input.Company,
		Location:    input.Location,
		StartDate:   input.StartDate,
		EndDate:     input.EndDate,
		UserID:      userIDUint,
	}

	result := config.DB.Create(&experience)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(201, gin.H{"data": experience})
}

func GetAllExperiences(c *gin.Context) {
	var experiences []models.Experience

	result := config.DB.Find(&experiences)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if len(experiences) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No experiences found"})
		return
	}

	c.JSON(http.StatusOK, experiences)
}

func GetExperienceByID(c *gin.Context) {
	id := c.Param("id")

	var experience models.Experience

	if err := config.DB.Where("ID = ?", id).First(&experience).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Experience not found"})
		return
	}

	c.JSON(http.StatusOK, experience)
}

func DeleteExperience(c *gin.Context) {
	id := c.Param("id")

	var experience models.Experience

	if err := config.DB.Where("ID = ?", id).First(&experience).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Experience not found"})
		return
	}

	config.DB.Delete(&experience)

	c.JSON(http.StatusOK, gin.H{"message": "Experience deleted successfully"})

}

func UpdateExperience(c *gin.Context) {
	id := c.Param("id")
	var input EditExperienceInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data", "details": err.Error()})
		return
	}

	var experience models.Experience

	if err := config.DB.Where("ID = ?", id).First(&experience).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Experience not found"})
		return
	}

	if input.Title != "" {
		experience.Title = input.Title
	}
	if input.Description != "" {
		experience.Description = input.Description
	}
	if input.Company != "" {
		experience.Company = input.Company
	}
	if input.Location != "" {
		experience.Location = input.Location
	}
	if input.StartDate != "" {
		experience.StartDate = input.StartDate
	}
	if input.EndDate != "" {
		experience.EndDate = input.EndDate
	}

	if err := config.DB.Save(&experience).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update experience"})
		return
	}

	c.JSON(http.StatusOK, experience)
}
