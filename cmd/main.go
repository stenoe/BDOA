package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stenoe/go-database/database"
)

func main() {
	database.ConnectDB()
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3002")
}
