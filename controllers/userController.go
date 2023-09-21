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
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Bad request", "error": err})
	}

	db.DB.Select("id,email").Where("email = ?", user.Email).First(&user)

	if user.ID != uuid.Nil {
		return c.Status(400).JSON(fiber.Map{"message": "User already exists"})
	}

	user.ID = uuid.New()

	//create user
	db.DB.Create(&user)

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    user,
	})
}

func GetUser(c *fiber.Ctx) error {
	var user models.User
	db.DB.Select("id,name,email,password,phone,address,profile_pic, rut").Where("email = ?", c.Params("email")).Find(&user).First(&user)
	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"message": "User not found"})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    user,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Bad request", "error": err})
	}

	db.DB.Select("id,name,email,password,phone,address,profile_pic, rut").Where("email = ?", c.Params("email")).Find(&user).First(&user)
	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"message": "User not found"})
	}
	//! falta completar el codigo
	// db.DB.Model(&user).Updates(models.User{
	// 	Name:       user.Name,
	// 	Email:      user.Email,
	// 	Password:   user.Password,
	// 	Phone:      user.Phone,
	// 	Address:    user.Address,
	// 	ProfilePic: user.ProfilePic,
	// 	Rut:        user.Rut,
	// })

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    user,
	})
}
