package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	userGroup := app.Group("/users")

	userGroup.Post("/", controllers.CreateUser)
	userGroup.Post("/verifyCode", controllers.CompareHashedCode)
	userGroup.Post("/resetPassword/:email", controllers.SendEmailVerification)

	userGroup.Get("/", controllers.GetUsers)
	//userGroup.Get("/:email", controllers.GetUserByEmail)
	userGroup.Get("/:id", controllers.GetUserById)

	userGroup.Put("/changePassword/", controllers.UpdatePassword)
	userGroup.Put("/:id", controllers.UpdateUser)

	authGroup := app.Group("/auth")

	authGroup.Post("/signUp", controllers.SignUp)
	authGroup.Post("/signIn", controllers.SignIn)

	contactGroup := app.Group("/contacts")

	contactGroup.Get("/", controllers.GetContacts)
	contactGroup.Get("/:user_id", controllers.GetContactsByUser)
	contactGroup.Post("/", controllers.CreateContact)

	messageGroup := app.Group("/messages")

	messageGroup.Get("/", controllers.GetMessages)
	messageGroup.Get("/:travel_route_id", controllers.GetMessages)
	messageGroup.Post("/", controllers.CreateMesssage)

	pictureGroup := app.Group("/pictures")

	pictureGroup.Get("/", controllers.GetPictures)
	pictureGroup.Get("/:picture_id", controllers.GetPictures)
	pictureGroup.Post("/", controllers.CreatePicture)

	travelRouteGroup := app.Group("/travel-routes")

	travelRouteGroup.Get("/", controllers.GetTravelRoutes)
	travelRouteGroup.Get("/:user_id", controllers.GetTravelRoutes)
	travelRouteGroup.Post("/", controllers.CreateTravelRoute)

}
