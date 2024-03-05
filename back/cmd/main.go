package main

import (
	"github.com/Go-nine9/go-nine9/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.ConnectDB()

	app := fiber.New()

	app.Use(cors.New())

	setupRoutes(app)
	app.Listen(":8097")
}
