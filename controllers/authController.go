package controllers

import (
	db "backend/config"
	"backend/models"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func SignUp(c *fiber.Ctx) error {

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad request",
			"error":   err})
	}

	return CreateUser(c)
}

func SignIn(c *fiber.Ctx) error {

	var user models.User
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad request",
			"error":   err})
	}
	//search user
	db.DB.Where("email = ?", body.Email).First(&user)
	if user.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found"})
	}

	//verify email
	if user.Email != body.Email {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Wrong email"})
	}

	//verify password
	passwordGood := VerifyPassword(body.Password, user.Password)
	fmt.Println("password verify : ", passwordGood)
	if !passwordGood {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Wrong password"})
	}

	//creo el jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"own": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	//lo firmo
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Token Expired or invalid",
			"success": false})
	}

	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: time.Now().Add(time.Hour * 24),
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Signed in successfully",
		"data":    user,
	})
}
