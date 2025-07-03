package routes

import (
	"personal-portfolio-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(r *gin.Engine) {
	authRoute := r.Group("/auth")
	{
		authRoute.POST("/register", controllers.Register)
		// authRoute.POST("/login", Login)
		// authRoute.GET("/profile", AuthMiddleware(), GetProfile)
	}
}
