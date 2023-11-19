package controllers

import (
	db "backend/config"
	"backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetTravelRoutes(c *fiber.Ctx) error {

	userID := c.Params("user_id")

	var travelRoutes []models.TravelRoute
	db.DB.Select(("id,start_point,end_point,date,user_id")).Where("user_id = ?", userID).Find(&travelRoutes)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    travelRoutes,
	})

}

func CreateTravelRoute(c *fiber.Ctx) error {
	var travelRoute models.TravelRoute

	if err := c.BodyParser(&travelRoute); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Error while parsing",
		})
	}

	var user models.User
	if err := db.DB.Where("id = ?", travelRoute.UserID).First(&user).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "User not found",
		})
	}

	travelRoute.TravelRouteID = uuid.New()

	if err := db.DB.Create(&travelRoute).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Error while creating",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    travelRoute,
	})
}
