package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func About(app *fiber.App) {
	// Render index
	app.Get("/about", func(c *fiber.Ctx) error {
		return c.Render("about", fiber.Map{
			"Title":   "Tentang Kami - Maunguli",
			"PageCSS": "/static/css/page/about.css",
			"Year":    time.Now().Year(),
		})
	})
}
