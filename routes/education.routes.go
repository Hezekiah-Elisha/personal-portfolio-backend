package routes

import (
	"personal-portfolio-backend/controllers"
	"personal-portfolio-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupEducationRoutes(r *gin.Engine) {
	educationRoute := r.Group("/educations")
	{
		educationRoute.GET("/", controllers.GetAllEducations)
		// educationRoute.GET("/:id", GetEducationByID)
		educationRoute.POST("/", middlewares.AuthMiddleware(), controllers.CreateEducation)
	}
}
