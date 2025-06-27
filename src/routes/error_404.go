package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func Erorr404(app *fiber.App) {
	// Custom 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).Render("404", fiber.Map{
			"Title": "404 Not Found",
			"Year":  time.Now().Year(),
		})
	})
}
