package controllers

import (
	db "backend/config"
	"backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

// testear si se crea el contacto con el id del usuario
func CreateContact(c *fiber.Ctx) error {
	var contact models.Contact
	var user models.User
	if err := c.BodyParser(&contact); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Error while parsing",
		})
	}
	contact.ID = uuid.New()
	if user.ID == contact.UserID {
		db.DB.Create(&contact)

		return c.Status(201).JSON(fiber.Map{
			"success": true,
			"message": "success",
			"data":    contact,
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    contact,
	})
}
