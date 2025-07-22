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
		experienceRoute.GET("/:id", controllers.GetExperienceByID)
		experienceRoute.DELETE("/:id", middlewares.AuthMiddleware(), controllers.DeleteExperience)
		experienceRoute.POST("/", middlewares.AuthMiddleware(), controllers.CreateExperience)
		experienceRoute.PATCH("/:id", middlewares.AuthMiddleware(), controllers.UpdateExperience)
	}
}
