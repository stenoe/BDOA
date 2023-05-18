package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stenoe/go-database/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)

	app.Get("/listall", handlers.ListFacts)

	app.Post("/fact", handlers.CreateFact)
}
