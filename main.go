package main

import "github.com/gofiber/fiber/v2"

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

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, World")
	})

	app.Get("/todos", getTodos)
	app.Post("/add/todo", createTodo)

	app.Listen(":80")
}
