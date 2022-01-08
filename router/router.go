package router

import (
	"fiber-go-lang/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	api.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, World")
	})

	api.Get("/todos", handler.GetTodos)
	api.Get("/todos/:id", handler.GetTodo)
	api.Post("/add/todo", handler.CreateTodo)
	api.Delete("/delete/todo/:id", handler.DeleteTodo)
	api.Patch("/edit/todo/:id", handler.EditTodo)

	api.Post("/auth", handler.Auth)

	api.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})
}
