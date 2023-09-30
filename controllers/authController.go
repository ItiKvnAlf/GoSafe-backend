package controllers

import (
	"backend/models"

	"github.com/gofiber/fiber/v2"
)

func SignUp(c *fiber.Ctx) error {

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad request",
			"error":   err})
	}

	return c.SendString("Sign Up")
}
