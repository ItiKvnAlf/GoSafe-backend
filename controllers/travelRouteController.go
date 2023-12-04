package controllers

import (
	db "backend/config"
	"backend/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateTravelRoute(c *fiber.Ctx) error {
	var travel_route models.Travel_route

	if err := c.BodyParser(&travel_route); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": fiber.ErrBadRequest.Message,
		})
	}

	travel_route.ID = uuid.New()

	//set date now
	travel_route.Date = time.Now()

	travel_route.Pictures = []models.Picture{}
	travel_route.Message = []models.Message{}
	travel_route.Geolocation = []models.Geolocation{}

	var user models.User
	//verifico que el usuario exista

	userID := travel_route.UserID
	if err := db.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User not found",
		})
	}

	if err := db.DB.Create(&travel_route).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": fiber.ErrInternalServerError.Message,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    travel_route,
	})
}

func GetTravelRoutes(c *fiber.Ctx) error {

	var Travel_routes []models.Travel_route

	db.DB.Preload("User").Preload("Pictures").Preload("Message").Preload("Geolocation").Find(&Travel_routes)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    Travel_routes,
	})

}
