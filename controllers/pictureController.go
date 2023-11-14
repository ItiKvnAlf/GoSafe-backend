package controllers

import (
	db "backend/config"
	"backend/models"

	"github.com/gofiber/fiber/v2"
)

func GetPicturesTravel(c *fiber.Ctx) error {

	travelRouteID := c.Params("travel_route_id")

	var pictures []models.Picture
	db.DB.Select(("id,travel_route_id,image")).Where("travel_route_id = ?", travelRouteID).Find(&pictures)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    pictures,
	})

}

// testear si se crea la imagen con su ruta de viaje
func CreatePicture(c *fiber.Ctx) error {
	var picture models.Picture

	if err := c.BodyParser(&picture); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Error while parsing",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    picture,
	})
}
