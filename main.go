package main

import (
	"fiber-go-lang/config"
	"fiber-go-lang/handler"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	// app.Use(middleware.Logger())
	// app.Use(middleware.Recover())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, World")
	})

	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}

	app.Get("/todos", handler.GetTodos)
	app.Get("/todos/:id", handler.GetTodo)
	app.Post("/add/todo", handler.CreateTodo)
	app.Delete("/delete/todo/:id", handler.DeleteTodo)
	app.Patch("/edit/todo/:id", handler.EditTodo)

	app.Post("/auth", handler.Auth)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	log.Fatal(app.Listen(config.Config("PORT")))
}
