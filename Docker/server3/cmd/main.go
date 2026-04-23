package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/stenoe/server3/database"
)

func main() {
	database.ConnectDb()

	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
