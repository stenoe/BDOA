package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stenoe/server3/database"
	"github.com/stenoe/server3/models"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Hello from server3! :)")
}

func ListMeasurements(c *fiber.Ctx) error {
	var measurements []models.Measurement

	database.DB.Db.Find(&measurements)

	return c.Status(200).JSON(measurements)
}

func CreateMeasurement(c *fiber.Ctx) error {
	measurement := new(models.Measurement)
	if err := c.BodyParser(measurement); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&measurement)

	return c.Status(200).JSON(measurement)
}
