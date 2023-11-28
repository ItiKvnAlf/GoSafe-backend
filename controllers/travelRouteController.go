package controllers

import (
	db "backend/config"
	"backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetTravelRoutes(c *fiber.Ctx) error {

	userID := c.Params("user_id")

	var Travel_routes []models.Travel_route
	db.DB.Select(("id,start_point,end_point,date,user_id")).Where("user_id = ?", userID).Find(&Travel_routes)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    Travel_routes,
	})

}

func CreateTravelRoute(c *fiber.Ctx) error {
	var Travel_route models.Travel_route

	if err := c.BodyParser(&Travel_route); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Error while parsing",
		})
	}

	var user models.User
	if err := db.DB.Where("id = ?", Travel_route.UserID).First(&user).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "User not found",
		})
	}

	Travel_route.ID = uuid.New()

	if err := db.DB.Create(&Travel_route).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Error while creating",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    Travel_route,
	})
}
