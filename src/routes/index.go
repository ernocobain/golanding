package routes

import (
	"time"

	"github/dhikrama/go/utils"

	"github.com/gofiber/fiber/v2"
)

func Index(app *fiber.App) {
	logoImage := utils.GeneratePictureHTML(
		"/static/images/logo-maunguli.png",
		"Jasa Bangunan Maunguli",
		[]int{130, 260},
		"130px",
	)

	app.Get("/", func(c *fiber.Ctx) error {
		canonicalURL := c.BaseURL() + c.Path()
		return c.Render("index", fiber.Map{
			"Title":       "Maunguli - Jasa Tukang Profesional Terpercaya",
			"Description": "Mulai dari bangun rumah, renovasi total, hingga perbaikan kecil...",
			"Canonical":   canonicalURL, // <-- Gunakan variabel dinamis
			"Year":        time.Now().Year(),
			"Logo":        logoImage,
		}, "layouts/main")
	})
}
