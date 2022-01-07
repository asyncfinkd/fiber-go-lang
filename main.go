package main

import "github.com/gofiber/fiber/v2"

type Todo struct {
	Id        int
	text      string
	completed bool
}

func getTodos(ctx *fiber.Ctx) error {
	item := Todo{
		1,
		"Hello, World",
		false,
	}

	return ctx.Status(fiber.StatusOK).JSON(item)
}

var todo Todo

func createTodo(ctx *fiber.Ctx) error {
	body := new(Todo)
	err := ctx.BodyParser(body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}

	todo = Todo{
		Id:        body.Id,
		text:      body.text,
		completed: false,
	}

	return ctx.Status(fiber.StatusOK).JSON(todo)
}

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, World")
	})

	app.Get("/todos", getTodos)
	app.Post("add/todo", createTodo)

	app.Listen(":80")
}
