package routes

import (
	"personal-portfolio-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine) {
	userRoute := r.Group("/users")
	{
		userRoute.GET("/", controllers.GetAllUsers)
	}
}
