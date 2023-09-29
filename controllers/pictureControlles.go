package controllers

import (
	db "backend/config"
	"backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetPicturesTravel(c *fiber.Ctx) error {

	travelRouteID := c.Params("travel_route_id")

	var pictures []models.Picture
	db.DB.Select(("id, user_id, name, url")).Where("travel_route_id = ?", travelRouteID).Find(&pictures)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    pictures,
	})

}

func CreatePicture(c *fiber.Ctx) error {
	var picture models.Picture

	if err := c.BodyParser(&picture); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Error while parsing",
		})
	}

	var travelRoute models.TravelRoute
	if err := db.DB.Where("id = ?", picture.TravelRouteID).First(&travelRoute).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Travel Route not found",
		})
	}

	picture.ID = uuid.New()

	if err := db.DB.Create(&picture).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Error while creating",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    picture,
	})
}
