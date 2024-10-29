package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/tegveer-singh123/blog/database"
	"github.com/tegveer-singh123/blog/routes"
)

func main(){
	database.Connect()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	app := fiber.New()

	routes.Setup(app)
	
	app.Listen(":"+port)
}