package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func Privacy(app *fiber.App) {
	// Render index
	app.Get("/privacy", func(c *fiber.Ctx) error {
		return c.Render("privacy", fiber.Map{
			"Title": "Kebijakan privasi - Maunguli",
			"Year":  time.Now().Year(),
		})
	})
}
