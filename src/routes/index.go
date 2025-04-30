package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func Index(app *fiber.App) {
	// Render index
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Jasa tukang bangunan jakarta",
			"Year":  time.Now().Year(),
		}, "layouts/main")
	})
}
