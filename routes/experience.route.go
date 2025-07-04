package routes

import (
	"personal-portfolio-backend/controllers"
	"personal-portfolio-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupExperienceRoutes(r *gin.Engine) {
	experienceRoute := r.Group("/experiences")
	{
		experienceRoute.GET("/", controllers.GetAllExperiences)
		// experienceRoute.GET("/:id", GetExperienceByID)
		experienceRoute.POST("/", middlewares.AuthMiddleware(), controllers.CreateExperience)
	}
}
