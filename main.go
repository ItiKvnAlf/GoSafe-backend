package main

import (
	"fmt"
	"log"
	"os"

	db "backend/config"
	"backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Server starting...")
	godotenv.Load()
	fmt.Println("Loaded env variables...")

	// Connect to database
	db.Connect()

	//create fiber app
	app := fiber.New()
	app.Use(cors.New(
		cors.Config{
			AllowCredentials: true,
			AllowOrigins:     "http://localhost:5173",
		},
	))

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	portString := os.Getenv("APP_PORT")
	if portString == "" {
		log.Fatal("PORT is not set in .env file")
	}

	routes.Setup(app)
	app.Listen(":" + portString)
}
