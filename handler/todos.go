package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	Id        int
	Name      string
	Completed bool
}

type Message struct {
	Success bool
}

var todos = []*Todo{
	{Id: 1, Name: "test", Completed: false},
	{Id: 2, Name: "test2", Completed: true},
}

// @Summary Get all todos
// @Description Get all todos
// @Tags todo
// @Accept json
// @Produce json
// @Success 200 {array} Todo{}
// @Failure 503 {object} Message{}
// @Router /api/todos [get]
func GetTodos(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(todos)
}

func CreateTodo(ctx *fiber.Ctx) error {
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

func GetTodo(ctx *fiber.Ctx) error {
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

func DeleteTodo(ctx *fiber.Ctx) error {
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

func EditTodo(ctx *fiber.Ctx) error {
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
