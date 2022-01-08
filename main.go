package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct {
	Id        int
	Name      string
	Completed bool
}

type User struct {
	Email    string
	Password string
}

var todos = []*Todo{
	{Id: 1, Name: "test", Completed: false},
	{Id: 2, Name: "test2", Completed: true},
}

var users = []*User{
	{Email: "ns@gmail.com", Password: "123123123"},
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

	todo := &Todo{
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

func editTodo(ctx *fiber.Ctx) error {
	type request struct {
		Name      *string
		Completed *bool
	}

	paramsId := ctx.Params("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
	}

	var body request
	err = ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse body",
		})
	}

	var todo *Todo

	for _, t := range todos {
		if t.Id == id {
			todo = t
			break
		}
	}

	if todo == nil {
		return ctx.Status(fiber.StatusNotFound).JSON("...")
	}

	if body.Name != nil {
		todo.Name = *body.Name
	}

	if body.Completed != nil {
		todo.Completed = *body.Completed
	}

	return ctx.Status(fiber.StatusOK).JSON(todo)
}

func auth(ctx *fiber.Ctx) error {
	type request struct {
		Email    *string
		Password *string
	}

	type TReturn struct {
		Success      bool
		Message      string
		Access_Token string
	}

	var body request
	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse body",
		})
	}

	if !strings.Contains(*body.Email, "@") {
		return ctx.Status(fiber.StatusBadRequest).JSON("email address you entered doesn't contain '@' sign")
	}

	if len(*body.Password) < 6 {
		return ctx.Status(fiber.StatusBadRequest).JSON("password you entered is too short")
	}

	for _, t := range users {
		loginValidate := t.Email == *body.Email && t.Password == *body.Password
		if loginValidate {
			OReturn := &TReturn{
				Message:      "Congratulation, you logged succesfully",
				Success:      true,
				Access_Token: "___",
			}
			return ctx.Status(fiber.StatusOK).JSON(OReturn)
		}
	}

	OReturn := &TReturn{
		Message: "Credentials incorrect.",
		Success: false,
	}
	return ctx.Status(fiber.StatusBadRequest).JSON(OReturn)
}

func main() {
	app := fiber.New()

	// app.Use(middleware.Logger())
	// app.Use(middleware.Recover())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, World")
	})

	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}

	app.Get("/todos", getTodos)
	app.Get("/todos/:id", getTodo)
	app.Post("/add/todo", createTodo)
	app.Post("/auth", auth)
	app.Delete("/delete/todo/:id", deleteTodo)
	app.Patch("/edit/todo/:id", editTodo)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	log.Fatal(app.Listen(os.Getenv("PORT")))
}
