package controllers

import (
	db "backend/config"
	"backend/models"

	"github.com/gofiber/fiber/v2"
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
