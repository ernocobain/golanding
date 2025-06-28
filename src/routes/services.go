package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func Services(app *fiber.App) {
	app.Get("/services", servicesHandler)
}

func servicesHandler(c *fiber.Ctx) error {
	// Handler ini merender template 'services.html'
	return c.Render("services", fiber.Map{
		"Title":   "Layanan Kami - Maunguli",
		"Year":    time.Now().Year(),
		"PageCSS": "/static/css/page/services.css", // Path ke CSS khusus halaman ini
	})
}
