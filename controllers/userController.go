package controllers

import (
	db "backend/config"
	"backend/models"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/go-gomail/gomail"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// create user
func CreateUser(c *fiber.Ctx) error {

	var user models.User

	//verify json is correct
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "success",
		})
	}

	db.DB.Select("id,name,email,password,phone,address,profile_pic, rut").Where("email = ? or rut = ? or phone = ?", user.Email, user.Rut, user.Phone).First(&user)

	//verify that user does not already exist
	if user.ID != uuid.Nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": fiber.ErrBadRequest.Message,
		})
	}

	//set uuid
	user.ID = uuid.New()

	//set contacts
	user.Contacts = []models.Contact{}
	user.TravelRoutes = []models.Travel_route{}

	//encrypt password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": fiber.ErrBadRequest.Message,
		})
	}

	//set password hashed
	user.Password = string(hash)

	//create user
	db.DB.Create(&user)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    user,
	})
}

func GetUsers(c *fiber.Ctx) error {

	//falta relacionar contactos con una funcion
	var users []models.User

	//esperar ver resultados
	//db.DB.Select("id, name, email, password, phone, address, profile_pic, rut").Find(&users)

	db.DB.Preload("Contacts").Preload("TravelRoutes").Find(&users)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    users,
	})
}

func GetUserById(c *fiber.Ctx) error {
	var user models.User

	userID := c.Params("id")

	db.DB.Where("users.id = ?", userID).Joins("Contacts").Joins("TravelRoutes").First(&user)
	if user.ID == uuid.Nil || user.Email == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": fiber.ErrNotFound.Message,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": false,
		"message": "success",
		"data":    user,
	})
}

func GetUserByEmail(c *fiber.Ctx) error {
	var user models.User

	userEmail := c.Params("id")

	db.DB.Where("users.email = ?", userEmail).Joins("Contacts").Joins("TravelRoutes").First(&user)
	if user.ID == uuid.Nil || user.Email == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": fiber.ErrNotFound.Message,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": false,
		"message": "success",
		"data":    user,
	})
}

func UpdateUser(c *fiber.Ctx) error {

	//Consultar por ingreso invalido de datos
	userID := c.Params("id")

	var user models.User

	//error: user not found
	if err := db.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": fiber.ErrNotFound.Message,
		})
	}
	var updatedUser models.User

	fmt.Println("user:", user)

	//error parser
	if err := c.BodyParser(updatedUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": fiber.ErrBadRequest.Message,
		})
	}

	user.Name = updatedUser.Name

	//para actualizar el correo debiera verificarlo con su codigo similar ala contraseña
	//user.Email = updatedUser.Email

	hash, err := bcrypt.GenerateFromPassword([]byte(updatedUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fiber.ErrInternalServerError.Message})
	}

	user.Password = string(hash)

	user.Phone = updatedUser.Phone
	user.Address = updatedUser.Address
	user.ProfilePic = updatedUser.ProfilePic

	db.DB.Save(&user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Usuario actualizado exitosamente",
		"data":    user,
	})
}

func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func SendEmailVerification(c *fiber.Ctx) error {
	var user models.User

	userEmail := c.Params("email")

	result := db.DB.Where("email = ?", userEmail).First(&user)
	if result.Error != nil {
		return c.Status(204).JSON(fiber.Map{
			"success": false,
			"message": "User not found",
		})
	}

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	min := 100000
	max := 999999
	code := random.Intn(max-min+1) + min
	expiration := time.Now().Add(5 * time.Minute)

	nodeEmail := os.Getenv("GOMAIL_USER")

	m := gomail.NewMessage()
	m.SetHeader("From", "GoSafe <"+nodeEmail+">")
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", "Restablecimiento de contraseña")

	body := fmt.Sprintf(`Hola %s,<br><br>
	Has solicitado restablecer tu contraseña. Tu código de verificación es el siguiente:<br><br>
	%s<br><br>
	Este código caduca dentro de 5 minutos. Si no has solicitado restablecer tu contraseña, ignora este correo.<br><br>
	Gracias. GoSafe.`, user.Name, strconv.Itoa(code))

	m.SetBody("text/html", body)

	port, err := strconv.Atoi(os.Getenv("GOMAIL_PORT"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Error converting GOMAIL_PORT to integer",
			"error":   err,
		})
	}

	d := gomail.NewDialer(os.Getenv("GOMAIL_HOST"), port, os.Getenv("GOMAIL_USER"), os.Getenv("GOMAIL_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
		return c.Status(202).JSON(fiber.Map{
			"success": false,
			"message": "Failed to send email verification code",
			"error":   err,
		})
	}

	codeStr := strconv.Itoa(code)
	hashedCode, err := hashCode(codeStr)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Failed to hash verification code",
			"error":   err,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Sent successfully",
		"data":    hashedCode,
		"expires": expiration,
	})
}

func CompareHashedCode(c *fiber.Ctx) error {

	var body struct {
		Code       string    `json:"code"`
		HashedCode string    `json:"hashedCode"`
		Expires    time.Time `json:"expiration"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad request",
			"error":   err})
	}

	err := bcrypt.CompareHashAndPassword([]byte(body.HashedCode), []byte(body.Code))
	if err != nil {
		return c.Status(200).JSON(fiber.Map{
			"success": false,
			"message": "Code does not match",
			"error":   err,
		})
	}

	if time.Now().After(body.Expires) {
		return c.Status(200).JSON(fiber.Map{
			"success": false,
			"message": "Expired",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Code Verified",
	})
}

func UpdatePassword(c *fiber.Ctx) error {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
			"error":   err,
		})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error encrypting new password",
			"error":   err,
		})
	}

	var user models.User
	if err := db.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err == nil {
		return c.Status(200).JSON(fiber.Map{
			"message": "Same password",
		})
	}

	updatedUser := new(models.User)
	if err := c.BodyParser(updatedUser); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
			"error":   err,
		})
	}

	user.Password = string(hash)
	db.DB.Save(&user)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Password updated successfully",
		"data":    user,
	})
}

func hashCode(password string) (string, error) {
	hashedCode, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedCode), nil
}
