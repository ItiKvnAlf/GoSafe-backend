package controllers

import (
	db "backend/config"
	"backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetMessageTravel(c *fiber.Ctx) error {

	travelRouteID := c.Params("travel_route_id")

	var message []models.Message
	db.DB.Select(("id,user_id,travel_route_id,geolocation_id,last_picture,default_message")).Where("travel_route_id = ?", travelRouteID).Find(&message)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    message,
	})
}

func CreateMesssage(c *fiber.Ctx) error {
	var message models.Message

	if err := c.BodyParser(&message); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Error while parsing",
		})
	}

	message.ID = uuid.New()

	if err := db.DB.Create(&message).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Error while creating",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    message,
	})
}
