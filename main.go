package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	Id        int
	Name      string
	Completed bool
}

var todos = []Todo{
	{Id: 1, Name: "test", Completed: false},
	{Id: 2, Name: "test2", Completed: true},
}

func getTodos(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(todos)
}

func createTodo(ctx *fiber.Ctx) error {
	type request struct {
		Name string
	}

	var body request

	err := ctx.BodyParser(&body)

	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
	}

	todo := Todo{
		Id:        len(todos) + 1,
		Name:      body.Name,
		Completed: false,
	}

	todos = append(todos, todo)

	return ctx.Status(fiber.StatusCreated).JSON(todo)
}

func getTodo(ctx *fiber.Ctx) error {
	paramsId := ctx.Params("id")

	id, err := strconv.Atoi(paramsId)

	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
	}

	for _, todo := range todos {
		if todo.Id == id {
			return ctx.Status(fiber.StatusOK).JSON(todo)
		}
	}

	return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "system error",
	})
}

func deleteTodo(ctx *fiber.Ctx) error {
	paramsId := ctx.Params("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
	}

	for i, todo := range todos {
		if todo.Id == id {
			todos = append(todos[0:i], todos[i+1:]...)
			return ctx.Status(fiber.StatusNoContent).JSON("...")
		}
	}

	return ctx.Status(fiber.StatusNotFound).JSON("...")
}

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, World")
	})

	app.Get("/todos", getTodos)
	app.Post("/add/todo", createTodo)
	app.Get("/todos/:id", getTodo)
	app.Delete("/delete/todo/:id", deleteTodo)

	app.Listen(":80")
}
