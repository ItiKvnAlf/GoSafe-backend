package controllers

import (
	db "backend/config"
	"backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	db.DB.Select("id, name, email, password, phone, address, profile_pic, rut").Find(&users)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    users,
	})
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	user.ID = uuid.New()
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Bad request", "error": err})
	}
	db.DB.Create(&user)

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    user,
	})
}

func UpdateUser(c *fiber.Ctx) error {

	userID := c.Params("id")

	var user models.User
	if err := db.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Usuario no encontrado"})
	}

	updatedUser := new(models.User)
	if err := c.BodyParser(updatedUser); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Solicitud incorrecta", "error": err})
	}

	user.Name = updatedUser.Name
	user.Email = updatedUser.Email
	user.Password = updatedUser.Password
	user.Phone = updatedUser.Phone
	user.Address = updatedUser.Address
	user.ProfilePic = updatedUser.ProfilePic
	user.Rut = updatedUser.Rut

	db.DB.Save(&user)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Usuario actualizado exitosamente",
		"data":    user,
	})
}
