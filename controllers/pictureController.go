package controllers

import (
	db "backend/config"
	"backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetPictures(c *fiber.Ctx) error {

	PictureID := c.Params("picture_id")

	var pictures []models.Picture
	db.DB.Select(("id,image")).Where("picture_id = ?", PictureID).Find(&pictures)

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

	picture.PictureID = uuid.New()

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
