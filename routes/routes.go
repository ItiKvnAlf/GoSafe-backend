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

	contactGroup := app.Group("/contacts")

	contactGroup.Get("/", controllers.GetContacts)
	contactGroup.Get("/:user_id", controllers.GetContactUser)
	contactGroup.Post("/", controllers.CreateContact)

	messageGroup := app.Group("/messages")

	messageGroup.Get("/:travel_route_id", controllers.GetMessageTravel)
	messageGroup.Post("/", controllers.CreateMesssage)

	pictureGroup := app.Group("/pictures")

	pictureGroup.Get("/:travel_route_id", controllers.GetPicturesTravel)
	pictureGroup.Post("/", controllers.CreatePicture)

}
