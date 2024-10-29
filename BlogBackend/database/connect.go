package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tegveer-singh123/blog/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DSN")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}

	log.Println("Connected to the database")

	DB = database

	database.AutoMigrate(
		&models.User{},
		&models.Blog{},
	)
}
