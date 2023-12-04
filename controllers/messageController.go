package controllers

import (
	db "backend/config"
	"backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateMesssage(c *fiber.Ctx) error {
	var message models.Message

	if err := c.BodyParser(&message); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Error while parsing",
		})
	}

	travelRouteID := message.TravelRouteID
	if err := db.DB.Where("travel_routes.id = ?", travelRouteID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Travel route not found",
		})
	}
	message.ID = uuid.New()
	if err := db.DB.Create(&message).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error while creating",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    message,
	})
}

func GetMessages(c *fiber.Ctx) error {

	var message []models.Message
	db.DB.Find(&message)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    message,
	})
}

func GetMessageById(c *fiber.Ctx) error {

	var message models.Message

	id := c.Params("id")

	db.DB.Where("id=?", id).First(&message)
	if message.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Message not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    message,
	})
}

func GetMessageByTravelRouteId(c *fiber.Ctx) error {

	var message models.Message

	travel_route_id := c.Params("travel_route_id")

	db.DB.Where("travel_route_id=?", travel_route_id).First(&message)
	if message.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Message not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    message,
	})
}

func UpdateMessage(c *fiber.Ctx) error {

	messageID := c.Params("id")

	var message models.Message

	db.DB.Where("id = ?", messageID).First(&message)
	if message.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Message not found",
		})
	}

	if err := c.BodyParser(&message); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Error while parsing",
		})
	}
	db.DB.Save(&message)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    message,
	})
}
