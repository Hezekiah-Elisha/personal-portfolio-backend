package routes

import (
	"personal-portfolio-backend/controllers"
	"personal-portfolio-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupProjectRoutes(r *gin.Engine) {
	projectRoute := r.Group("/projects")
	{
		projectRoute.GET("/", controllers.GetAllProjects)
		// projectRoute.GET("/:id", GetProjectByID)
		projectRoute.POST("/", middlewares.AuthMiddleware(), controllers.CreateProject)
	}
}
