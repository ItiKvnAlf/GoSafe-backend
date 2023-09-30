package controllers

import (
	db "backend/config"
	"backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c *fiber.Ctx) error {

	//falta relacionar contactos con una funcion
	var users []models.User
	db.DB.Select("id, name, email, password, phone, address, profile_pic, rut").Find(&users)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    users,
	})
}

// create user
func CreateUser(c *fiber.Ctx) error {

	/*
		hacer test de que no se pueda crear un usuario con el mismo :
		   	email
		   	phone
		   	rut

		hacer test de que se pueda crear un usuario con el mismo :
			address
	*/
	var user models.User

	//verify json is correct
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad request",
			"error":   err})
	}

	db.DB.Select("id,name,email,password,phone,address,profile_pic, rut").Where("email = ? or rut = ? or phone = ?", user.Email, user.Rut, user.Phone).First(&user)

	//verify that user does not already exist
	if user.ID != uuid.Nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "User already registred"})
	}

	//set uuid
	user.ID = uuid.New()

	//set contacts
	user.Contacts = []models.Contact{}

	//encrypt password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error encrypting password",
			"error":   err})
	}

	//set password hashed
	user.Password = string(hash)

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

	userID := c.Params("id")

	db.DB.Select("id,name,email,password,phone,address,profile_pic, rut").Where("id = ?", userID).Find(&user).First(&user)
	if user.ID == uuid.Nil || user.Email == "" {
		return c.Status(404).JSON(fiber.Map{"message": "User not found"})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    user,
	})
}

func UpdateUser(c *fiber.Ctx) error {

	userID := c.Params("id")

	var user models.User
	if err := db.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found"})
	}

	updatedUser := new(models.User)
	if err := c.BodyParser(updatedUser); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
			"error":   err})
	}

	user.Name = updatedUser.Name
	user.Email = updatedUser.Email

	hash, err := bcrypt.GenerateFromPassword([]byte(updatedUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error encrypting password",
			"error":   err})
	}
	user.Password = string(hash)

	user.Phone = updatedUser.Phone
	user.Address = updatedUser.Address
	user.ProfilePic = updatedUser.ProfilePic

	db.DB.Save(&user)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Usuario actualizado exitosamente",
		"data":    user,
	})
}
