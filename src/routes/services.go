package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func Services(app *fiber.App) {
	// Render index
	app.Get("/services", func(c *fiber.Ctx) error {
		return c.Render("services", fiber.Map{
			"Title": "Semua layanan - Maunguli",
			"Year":  time.Now().Year(),
		})
	})
}
