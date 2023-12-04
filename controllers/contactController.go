package controllers

import (
	db "backend/config"
	"backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateContact(c *fiber.Ctx) error {
	var contact models.Contact

	if err := c.BodyParser(&contact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Error while parsing",
		})
	}

	userID := contact.UserID
	if err := db.DB.Where("users.id = ?", userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User not found",
		})
	}

	contact.ID = uuid.New()

	if err := db.DB.Create(&contact).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error while creating contact",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Contact created successfully",
		"data":    contact,
	})
}

func GetContacts(c *fiber.Ctx) error {
	var contacts []models.Contact
	db.DB.Find(&contacts)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    contacts,
	})

}

func GetContactsByUser(c *fiber.Ctx) error {

	userID := c.Params("user_id")

	var contacts []models.Contact
	db.DB.Where("user_id = ?", userID).Find(&contacts)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    contacts,
	})
}

func GetContactById(c *fiber.Ctx) error {

	ID := c.Params("id")

	var contacts []models.Contact
	db.DB.Where("id = ?", ID).Find(&contacts)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    contacts,
	})
}

func UpdateContact(c *fiber.Ctx) error {

	ID := c.Params("id")

	var contact models.Contact
	db.DB.Where("id = ?", ID).First(&contact)

	if contact.Name == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Contact not found",
		})
	}

	if err := c.BodyParser(&contact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Error while parsing",
		})
	}

	db.DB.Save(&contact)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Contact updated successfully",
		"data":    contact,
	})
}

func DeleteContact(c *fiber.Ctx) error {

	ID := c.Params("id")

	var contact models.Contact
	db.DB.Where("id = ?", ID).First(&contact)

	if contact.Name == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Contact not found",
		})
	}

	db.DB.Delete(&contact)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Contact deleted successfully",
	})
}
