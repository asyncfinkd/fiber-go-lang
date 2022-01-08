package main

import (
	"fiber-go-lang/config"
	"fiber-go-lang/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	router.SetupRoutes(app)

	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}

	log.Fatal(app.Listen(config.Config("PORT")))
}
