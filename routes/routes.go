package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	//rutas publicas

	authGroup := app.Group("/auth")
	authGroup.Post("/signUp", controllers.SignUp)
	authGroup.Post("/signIn", controllers.SignIn)

	userGroup := app.Group("/users")
	userGroup.Get("/email/:email", controllers.GetUserByEmail)

	//middleware global ... para desabilitarlo comentelo
	AuthMiddleware := middleware.AuthMiddleware()

	app.Use(AuthMiddleware)

	//rutas protegidas

	userGroup.Post("/", controllers.CreateUser)
	userGroup.Post("/verifyCode", controllers.CompareHashedCode)
	userGroup.Post("/resetPassword/:email", controllers.SendEmailVerification)
	userGroup.Get("/", controllers.GetUsers)
	userGroup.Get("/:id", controllers.GetUserById)
	userGroup.Put("/changePassword/", controllers.UpdatePassword)
	userGroup.Patch("/:id", controllers.UpdateUser)

	//falta delete user: revisar si es necesario

	//checks if the user is logged in
	contactGroup := app.Group("/contacts")
	contactGroup.Post("/", controllers.CreateContact)
	contactGroup.Get("/", controllers.GetContacts)
	contactGroup.Get("/:id", controllers.GetContactById)
	//contactGroup.Get("/:user_id", controllers.GetContactsByUser)
	contactGroup.Patch("/:id", controllers.UpdateContact)
	contactGroup.Delete("/:id", controllers.DeleteContact)

	//checks if the user is logged in
	travelRouteGroup := app.Group("/travel-routes")
	travelRouteGroup.Get("/", controllers.GetTravelRoutes)
	travelRouteGroup.Get("/:id", controllers.GetTravelRoutesById)
	travelRouteGroup.Post("/", controllers.CreateTravelRoute)
	travelRouteGroup.Patch("/:id", controllers.UpdateTravelRoute)
	travelRouteGroup.Delete("/:id", controllers.DeleteTravelRoute)

	//checks if the user is logged in
	geolocationGroup := app.Group("/geolocations")
	geolocationGroup.Post("/", controllers.CreateGeolocation)
	geolocationGroup.Get("/", controllers.GetGeolocations)
	geolocationGroup.Get("/:id", controllers.GetGeolocationById)
	geolocationGroup.Patch("/:id", controllers.UpdateGeolocation)
	geolocationGroup.Delete("/:id", controllers.DeleteGeolocation)

	//checks if the user is logged in
	messageGroup := app.Group("/messages")
	messageGroup.Post("/", controllers.CreateMesssage)
	messageGroup.Get("/", controllers.GetMessages)
	messageGroup.Get("/:id", controllers.GetMessageById)
	messageGroup.Patch("/:id", controllers.UpdateMessage)
	messageGroup.Delete("/:id", controllers.DeleteMessage)

	//checks if the user is logged in
	pictureGroup := app.Group("/pictures")
	pictureGroup.Post("/", controllers.CreatePicture)
	pictureGroup.Get("/", controllers.GetPictures)
	pictureGroup.Get("/:id", controllers.GetPictureById)
	pictureGroup.Patch("/:id", controllers.UpdatePicture)
	pictureGroup.Delete("/:id", controllers.DeletePicture)
}
