package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func Contact(app *fiber.App) {
	// Render index
	app.Get("/contact", func(c *fiber.Ctx) error {
		return c.Render("contact", fiber.Map{
			"Title":   "Kontak Kami - Maunguli",
			"PageCSS": "/static/css/page/contact.css",
			"Year":    time.Now().Year(),
		})
	})
}
