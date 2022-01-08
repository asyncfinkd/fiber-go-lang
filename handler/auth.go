package handler

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Email    string
	Password string
}

var users = []*User{
	{Email: "ns@gmail.com", Password: "123123123"},
}

func Auth(ctx *fiber.Ctx) error {
	type request struct {
		Email    *string
		Password *string
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
			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
				"success":      true,
				"message":      "Congratulation, you logged succesfully",
				"access_token": "___",
			})
		}
	}

	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"success": false,
		"message": "credentials incorrect.",
	})
}
