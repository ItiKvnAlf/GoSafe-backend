package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	userGroup := app.Group("/users")

	userGroup.Get("/", controllers.GetUsers)
	userGroup.Post("/", controllers.CreateUser)
	userGroup.Get("/:id", controllers.GetUser)
	userGroup.Put("/:id", controllers.UpdateUser)

	authGroup := app.Group("/auth")

	authGroup.Post("/signup", controllers.SignUp)
}
