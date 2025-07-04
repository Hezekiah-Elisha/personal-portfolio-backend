package routes

import (
	"personal-portfolio-backend/controllers"
	"personal-portfolio-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine) {
	userRoute := r.Group("/users")
	{
		userRoute.GET("/", controllers.GetAllUsers)
		userRoute.POST("/", controllers.Createuser)
		userRoute.GET("/:email", controllers.FindUserByEmail)
		userRoute.DELETE("/:id", middlewares.AuthMiddleware(), controllers.DeleteUser)
	}
}
