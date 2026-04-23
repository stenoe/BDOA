package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/stenoe/server3/database"
	"github.com/stenoe/server3/models"
)

func Home(c fiber.Ctx) error {
	return c.SendString("Hello, from server3 now with air!")
}

func ListMeasurements(c fiber.Ctx) error {
	var measurements []models.Measurement

	result := database.DB.Db.Find(&measurements)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve measurements",
		})
	}

	return c.Status(fiber.StatusOK).JSON(measurements)
}

func CreateMeasurement(c fiber.Ctx) error {
	measurement := new(models.Measurement)
	if err := c.Bind().JSON(measurement); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	result := database.DB.Db.Create(&measurement)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create measurement",
		})
	}
	return c.Status(fiber.StatusOK).JSON(measurement)

}
