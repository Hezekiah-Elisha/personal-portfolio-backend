package controllers

import (
	"net/http"
	"personal-portfolio-backend/config"
	"personal-portfolio-backend/models"
	"time"

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

func GetEducationByID(c *gin.Context) {
	id := c.Param("id")
	var education models.Education
	if err := config.DB.First(&education, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Education not found"})
		return
	}

	educationResponse := EducationResponse{
		ID:           education.ID,
		Institution:  education.Institution,
		Degree:       education.Degree,
		FieldOfStudy: education.FieldOfStudy,
		StartDate:    education.StartDate.Format("2006-01-02"),
		EndDate:      education.EndDate.Format("2006-01-02"),
		Skills:       education.Skills,
		UserID:       education.UserID,
		CreatedAt:    education.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    education.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	c.JSON(http.StatusOK, educationResponse)
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

	startDate, err := time.Parse("2006-01-02", input.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format. Use YYYY-MM-DD"})
		return
	}
	var endDate time.Time
	if input.EndDate != "" {
		endDate, err = time.Parse("2006-01-02", input.EndDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date format. Use YYYY-MM-DD"})
			return
		}
	}

	education := models.Education{
		Institution:  input.Institution,
		Degree:       input.Degree,
		FieldOfStudy: input.FieldOfStudy,
		StartDate:    startDate,
		EndDate:      endDate,
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
		StartDate:    education.StartDate.Format("2006-01-02"),
		EndDate:      education.EndDate.Format("2006-01-02"),
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
			StartDate:    education.StartDate.Format("2006-01-02"),
			EndDate:      education.EndDate.Format("2006-01-02"),
			Skills:       education.Skills,
			UserID:       education.UserID,
			CreatedAt:    education.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:    education.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
		educationResponses = append(educationResponses, educationResponse)
	}

	c.JSON(http.StatusOK, educationResponses)
}

func DeleteEducation(c *gin.Context) {
	id := c.Param("id")
	var education models.Education

	if err := config.DB.First(&education, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Education not found"})
		return
	}

	if err := config.DB.Delete(&education).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete education"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Education deleted successfully"})
}

func UpdateEducation(c *gin.Context) {
	id := c.Param("id")
	var education models.Education

	if err := config.DB.First(&education, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Education not found"})
		return
	}

	var input CreateEducationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data", "details": err.Error()})
		return
	}

	startDate, err := time.Parse("2006-01-02", input.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format. Use YYYY-MM-DD"})
		return
	}
	var endDate time.Time
	if input.EndDate != "" {
		endDate, err = time.Parse("2006-01-02", input.EndDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date format. Use YYYY-MM-DD"})
			return
		}
	}

	education.Institution = input.Institution
	education.Degree = input.Degree
	education.FieldOfStudy = input.FieldOfStudy
	education.StartDate = startDate
	education.EndDate = endDate
	education.Skills = input.Skills

	if err := config.DB.Save(&education).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update education"})
		return
	}

	educationResponse := EducationResponse{
		ID:           education.ID,
		Institution:  education.Institution,
		Degree:       education.Degree,
		FieldOfStudy: education.FieldOfStudy,
		StartDate:    education.StartDate.Format("2006-01-02"),
		EndDate:      education.EndDate.Format("2006-01-02"),
		Skills:       education.Skills,
		UserID:       education.UserID,
		CreatedAt:    education.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    education.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	c.JSON(http.StatusOK, educationResponse)
}
