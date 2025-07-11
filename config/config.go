package config

import (
	"log"
	"os"
	"personal-portfolio-backend/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
		// Don't return here in production, env vars might be set elsewhere
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Add debug logging to check if env vars are set
	log.Printf("Connecting to database: %s@%s:%s/%s", dbUser, dbHost, dbPort, dbName)

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// CHECK ERROR IMMEDIATELY AFTER gorm.Open()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
		return nil
	}

	// Now it's safe to use the db connection
	if os.Getenv("GIN_MODE") == "debug" {
		err = db.Debug().AutoMigrate(&models.User{}, &models.Experience{}, &models.Project{}, &models.Education{})
		if err != nil {
			log.Fatalf("Error running database migrations: %v", err)
			return nil
		}
	} else {
		err = db.AutoMigrate(&models.User{}, &models.Experience{}, &models.Project{}, &models.Education{})
		if err != nil {
			log.Fatalf("Error running database migrations: %v", err)
			return nil
		}
	}

	DB = db
	log.Println("Database connection established successfully")

	return DB
}
