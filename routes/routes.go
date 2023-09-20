package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	userGroup := app.Group("/users")

	userGroup.Get("/", controllers.GetUsers)
	userGroup.Post("/", controllers.CreateUser)
	userGroup.Get("/:email", controllers.GetUser)

}
