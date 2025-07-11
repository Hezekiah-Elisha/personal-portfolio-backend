package main

import (
	"log"
	"os"
	"personal-portfolio-backend/config"
	"personal-portfolio-backend/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)

	ginMode := os.Getenv("GIN_MODE")

	switch ginMode {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
	// Handle database connection errors
	db := config.ConnectDB()
	if db == nil {
		log.Fatal("Failed to connect to database")
		return
	}

	r := gin.Default()
	r.Use(gin.Logger())

	allowedOrigins := []string{"http://localhost:3000", "http://127.0.0.1:3000"}
	if ginMode == "release" {
		// Add your production domain
		allowedOrigins = append(allowedOrigins, "https://hezekiah.hub.ke", "https://api.hezekiah.hub.ke")
	}

	r.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Personal Portfolio Backend API",
		})
	})
	routes.SetupAuthRoutes(r)
	routes.SetupUserRoutes(r)
	routes.SetupExperienceRoutes(r)
	routes.SetupEducationRoutes(r)
	routes.SetupProjectRoutes(r)

	r.Run(":8081")
}
