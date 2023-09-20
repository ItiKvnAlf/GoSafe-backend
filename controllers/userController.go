package controllers

import (
	db "backend/config"
	"backend/models"
	"fmt"

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

	db.DB.Select("email").Where("email = ?", user.Email).First(&user)
	fmt.Println("user :", user.Email)

	// if user ==  {
	// 	return c.Status(400).JSON(fiber.Map{"message": "User already exists"})
	// }

	//db.DB.Create(&user)

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    user,
	})
}

func GetUser(c *fiber.Ctx) error {
	var user models.User
	db.DB.Select("id,name,email,password,phone,address,profile_pic, rut").Where("email = ?", c.Params("email")).Find(&user).First(&user)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    user,
	})
}
