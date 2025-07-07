package controllers

import (
	"net/http"
	"personal-portfolio-backend/config"
	"personal-portfolio-backend/models"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateProjectInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	TechStack   string `json:"tech_stack" binding:"required"`
	Link        string `json:"link"`
	SourceCode  string `json:"source_code" binding:"required"`
	Image       string `json:"image" binding:"required"`
	UserID      uint   `json:"user_id"`
}

type ProjectResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	TechStack   string `json:"tech_stack"`
	Link        string `json:"link"`
	SourceCode  string `json:"source_code"`
	Image       string `json:"image"`
	UserID      uint   `json:"user_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func CreateProject(c *gin.Context) {
	var input CreateProjectInput

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

	project := models.Project{
		Title:       input.Title,
		Description: input.Description,
		TechStack:   input.TechStack,
		Link:        stringPtr(input.Link),
		SourceCode:  stringPtr(input.SourceCode),
		Image:       input.Image,
		UserID:      userIDUint,
	}
	if err := config.DB.Create(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, ProjectResponse{
		ID:          project.ID,
		Title:       project.Title,
		Description: project.Description,
		TechStack:   project.TechStack,
		Link:        derefString(project.Link),
		SourceCode:  derefString(project.SourceCode),
		Image:       project.Image,
		UserID:      project.UserID,
		CreatedAt:   project.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   project.UpdatedAt.Format(time.RFC3339),
	})
}

func derefString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func stringPtr(s string) *string {
	return &s
}

func GetAllProjects(c *gin.Context) {
	var projects []models.Project
	if err := config.DB.Find(&projects).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve projects", "details": err.Error()})
		return
	}

	var projectResponses []ProjectResponse
	for _, project := range projects {
		projectResponses = append(projectResponses, ProjectResponse{
			ID:          project.ID,
			Title:       project.Title,
			Description: project.Description,
			TechStack:   project.TechStack,
			Link:        derefString(project.Link),
			SourceCode:  derefString(project.SourceCode),
			Image:       project.Image,
			UserID:      project.UserID,
			CreatedAt:   project.CreatedAt.String(),
			UpdatedAt:   project.UpdatedAt.String(),
		})
	}

	c.JSON(http.StatusOK, projectResponses)
}
