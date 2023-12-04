package controllers

import (
	db "backend/config"
	"backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateGeolocation(c *fiber.Ctx) error {
	var geolocation models.Geolocation

	if err := c.BodyParser(&geolocation); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Error while parsing",
		})
	}

	travelRouteID := geolocation.TravelRouteID
	if err := db.DB.Where("travel_routes.id = ?", travelRouteID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Travel route not found",
		})
	}

	geolocation.ID = uuid.New()
	if err := db.DB.Create(&geolocation).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error while creating",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    geolocation,
	})
}

func GetGeolocations(c *fiber.Ctx) error {

	var geolocations []models.Geolocation
	db.DB.Find(&geolocations)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    geolocations,
	})
}

func GetGeolocationById(c *fiber.Ctx) error {

	ID := c.Params("id")

	var geolocations []models.Geolocation
	db.DB.Where("id = ?", ID).Find(&geolocations)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    geolocations,
	})
}

func GetGeolocationByTravelRouteId(c *fiber.Ctx) error {

	ID := c.Params("travel-route-id")

	var geolocations []models.Geolocation
	db.DB.Where("travel_route_id = ?", ID).Find(&geolocations)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    geolocations,
	})
}

func UpdateGeolocation(c *fiber.Ctx) error {

	ID := c.Params("id")

	var geolocation models.Geolocation
	db.DB.Where("id = ?", ID).First(&geolocation)

	if geolocation.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Geolocation not found",
		})
	}

	if err := c.BodyParser(&geolocation); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Error while parsing",
		})
	}
	db.DB.Save(&geolocation)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    geolocation,
	})
}

func DeleteGeolocation(c *fiber.Ctx) error {

	ID := c.Params("id")

	var geolocation models.Geolocation
	db.DB.Where("id = ?", ID).First(&geolocation)

	if geolocation.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Geolocation not found",
		})
	}

	db.DB.Delete(&geolocation)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
	})
}
