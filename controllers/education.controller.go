package controllers

import (
	"net/http"
	"personal-portfolio-backend/config"
	"personal-portfolio-backend/models"

	"github.com/gin-gonic/gin"
)

type CreateEducationInput struct {
	Institution  string `json:"institution" binding:"required"`
	Degree       string `json:"degree" binding:"required"`
	FieldOfStudy string `json:"field_of_study" binding:"required"`
	StartDate    string `json:"start_date" binding:"required"`
	EndDate      string `json:"end_date"`
	Skills       string `json:"skills" binding:"required"`
	UserID       uint   `json:"user_id"`
}
type EducationResponse struct {
	ID           uint   `json:"id"`
	Institution  string `json:"institution"`
	Degree       string `json:"degree"`
	FieldOfStudy string `json:"field_of_study"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
	Skills       string `json:"skills"`
	UserID       uint   `json:"user_id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

func CreateEducation(c *gin.Context) {
	var input CreateEducationInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data", "details": err.Error()})
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	userIDUint, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	education := models.Education{
		Institution:  input.Institution,
		Degree:       input.Degree,
		FieldOfStudy: input.FieldOfStudy,
		StartDate:    input.StartDate,
		EndDate:      input.EndDate,
		Skills:       input.Skills,
		UserID:       userIDUint,
	}
	if err := config.DB.Create(&education).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	educationResponse := EducationResponse{
		ID:           education.ID,
		Institution:  education.Institution,
		Degree:       education.Degree,
		FieldOfStudy: education.FieldOfStudy,
		StartDate:    education.StartDate,
		EndDate:      education.EndDate,
		Skills:       education.Skills,
		UserID:       education.UserID,
		CreatedAt:    education.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    education.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	c.JSON(http.StatusCreated, educationResponse)
}

func GetAllEducations(c *gin.Context) {
	var educations []models.Education
	if err := config.DB.Find(&educations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(educations) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No educations found"})
		return
	}

	var educationResponses []EducationResponse
	for _, education := range educations {
		educationResponse := EducationResponse{
			ID:           education.ID,
			Institution:  education.Institution,
			Degree:       education.Degree,
			FieldOfStudy: education.FieldOfStudy,
			StartDate:    education.StartDate,
			EndDate:      education.EndDate,
			Skills:       education.Skills,
			UserID:       education.UserID,
			CreatedAt:    education.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:    education.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
		educationResponses = append(educationResponses, educationResponse)
	}

	c.JSON(http.StatusOK, educationResponses)
}
