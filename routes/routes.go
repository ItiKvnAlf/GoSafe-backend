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

	//checks if the user is logged in
	travelRouteGroup := app.Group("/travel-routes")
	travelRouteGroup.Get("/", controllers.GetTravelRoutes)
	travelRouteGroup.Get("/:id", controllers.GetTravelRoutesById)
	travelRouteGroup.Post("/", controllers.CreateTravelRoute)
	travelRouteGroup.Patch("/:id", controllers.UpdateTravelRoute)
	travelRouteGroup.Delete("/:id", controllers.DeleteTravelRoute)

	authGroup := app.Group("/auth")
	authGroup.Post("/signUp", controllers.SignUp)
	authGroup.Post("/signIn", controllers.SignIn)

	contactGroup := app.Group("/contacts")
	contactGroup.Get("/", controllers.GetContacts)
	contactGroup.Get("/:user_id", controllers.GetContactsByUser)
	contactGroup.Post("/", controllers.CreateContact)

	//checks if the user is logged in
	messageGroup := app.Group("/messages")
	messageGroup.Post("/", controllers.CreateMesssage)
	messageGroup.Get("/", controllers.GetMessages)
	messageGroup.Get("/:id", controllers.GetMessageById)
	messageGroup.Patch("/:id", controllers.UpdateMessage)
	messageGroup.Delete("/:id", controllers.DeleteMessage)

	pictureGroup := app.Group("/pictures")
	pictureGroup.Get("/", controllers.GetPictures)
	pictureGroup.Get("/:picture_id", controllers.GetPictures)
	pictureGroup.Post("/", controllers.CreatePicture)

}
