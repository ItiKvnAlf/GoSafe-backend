package controllers

import (
	db "backend/config"
	"backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetGeolocations(c *fiber.Ctx) error {

	ID := c.Params("messageID")

	var geolocations []models.Geolocation
	db.DB.Select(("id,default_message,travel_route_id")).Where("messageID = ?", ID).Find(&geolocations)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    geolocations,
	})
}

func CreateGeolocation(c *fiber.Ctx) error {
	var geolocation models.Geolocation

	if err := c.BodyParser(&geolocation); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Error while parsing",
		})
	}

	geolocation.ID = uuid.New()
	if err := db.DB.Create(&geolocation).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Error while creating",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    geolocation,
	})
}
