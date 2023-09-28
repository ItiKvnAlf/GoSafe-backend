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

	if err := c.BodyParser(&contact); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Error while parsing",
		})
	}

	var user models.User
	if err := db.DB.Where("id = ?", contact.UserID).First(&user).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "User not found",
		})
	}

	contact.ID = uuid.New()

	if err := db.DB.Create(&contact).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Error while creating contact",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "Contact created successfully",
		"data":    contact,
	})
}
