package controllers

import (
	db "backend/config"
	"backend/models"

	"github.com/gofiber/fiber/v2"
)

func GetContacts(c *fiber.Ctx) error {
	var contacts []models.Contact
	db.DB.Select(("id, user_id, name, email, phone")).Find(&contacts)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    contacts,
	})

}
