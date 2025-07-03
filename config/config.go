package config

import (
	"log"
	"os"
	"personal-portfolio-backend/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	models.User
}

type Experience struct {
	gorm.Model
	models.Experience
}

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
		return nil
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&User{}, &Experience{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
		// return nil
	}
	DB = db
	log.Println("Database connection established successfully")

	return DB
}
