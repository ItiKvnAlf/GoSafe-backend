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

func GetTravelRoutesById(c *fiber.Ctx) error {

	var travel_routes models.Travel_route

	travelRouteID := c.Params("id")

	db.DB.Preload("User").Preload("Pictures").Preload("Message").Preload("Geolocation").Where("travel_routes.id=?", travelRouteID).Find(&travel_routes)
	if travel_routes.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Travel route not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    travel_routes,
	})

}

func GetTravelRoutesByUserID(c *fiber.Ctx) error {

	var travel_routes models.Travel_route

	userID := c.Params("id")

	db.DB.Preload("User").Preload("Pictures").Preload("Message").Preload("Geolocation").Where("travel_routes.user_id=?", userID).Find(&travel_routes)
	if travel_routes.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Travel route not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    travel_routes,
	})

}
func UpdateTravelRoute(c *fiber.Ctx) error {

	travelRouteID := c.Params("id")

	var travel_route models.Travel_route

	db.DB.Where("id = ?", travelRouteID).First(&travel_route)
	if travel_route.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Travel route not found",
		})
	}

	//solo puede actualizar dos campos: start_point y end_point

	if err := c.BodyParser(&travel_route); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": fiber.ErrBadRequest.Message,
		})
	}

	db.DB.Save(&travel_route)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    travel_route,
	})

}

func DeleteTravelRoute(c *fiber.Ctx) error {

	travelRouteID := c.Params("id")

	var travel_route models.Travel_route

	db.DB.Where("id = ?", travelRouteID).First(&travel_route)
	if travel_route.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Travel route not found",
		})
	}

	//actualizo la ruta en el usuario
	var user models.User

	user_id := travel_route.UserID
	db.DB.Where("users.id = ?", user_id).First(&user)

	//elimino la ruta de la lista de rutas del usuario
	for i, route := range user.TravelRoutes {
		if route.ID == travel_route.ID {
			user.TravelRoutes = append(user.TravelRoutes[:i], user.TravelRoutes[i+1:]...)
			break
		}
	}

	db.DB.Save(&user)

	db.DB.Delete(&travel_route)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "success",
	})

}
