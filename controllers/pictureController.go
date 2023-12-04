package controllers

import (
	db "backend/config"
	"backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreatePicture(c *fiber.Ctx) error {
	var picture models.Picture

	if err := c.BodyParser(&picture); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Error while parsing",
		})
	}

	picture.ID = uuid.New()

	if err := db.DB.Create(&picture).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error while creating",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    picture,
	})
}

func GetPictures(c *fiber.Ctx) error {

	var pictures []models.Picture
	db.DB.Find(&pictures)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    pictures,
	})

}
func GetPictureById(c *fiber.Ctx) error {

	PictureID := c.Params("picture_id")

	var pictures []models.Picture
	db.DB.Where("id = ?", PictureID).Find(&pictures)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    pictures,
	})

}

func GetPictureByTravelRouteId(c *fiber.Ctx) error {

	var picture models.Picture

	travel_route_id := c.Params("travel_route_id")

	db.DB.Where("travel_route_id=?", travel_route_id).First(&picture)
	if picture.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Message not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    picture,
	})
}

func UpdatePicture(c *fiber.Ctx) error {

	var picture models.Picture

	id := c.Params("id")

	if err := db.DB.Where("id = ?", id).First(&picture).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Picture not found",
		})
	}

	if err := c.BodyParser(&picture); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Error while parsing",
		})
	}

	db.DB.Save(&picture)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    picture,
	})
}

func DeletePicture(c *fiber.Ctx) error {

	id := c.Params("id")

	var picture models.Picture

	db.DB.Where("id = ?", id).First(&picture)
	if picture.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Picture not found",
		})
	}

	db.DB.Delete(&picture)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
	})
}
