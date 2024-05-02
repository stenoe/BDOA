package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stenoe/server3/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)

	app.Get("/measurement", handlers.ListMeasurements)

	app.Post("/measurement", handlers.CreateMeasurement)
}
