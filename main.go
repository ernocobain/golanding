package main

import (
	"log"

	m "github/dhikrama/go/src"
	r "github/dhikrama/go/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Create a new engine
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(helmet.New(helmet.Config{
		HSTSMaxAge:                300,
		HSTSExcludeSubdomains:     true,
		XSSProtection:             "0",
		XFrameOptions:             "SAMEORIGIN",
		ContentSecurityPolicy:     "same-origin",
		CrossOriginOpenerPolicy:   "same-origin",
		CrossOriginEmbedderPolicy: "require-corp",
	}))
	app.Static("/", "./static")

	r.Index(app)
	m.FormContact(app)
	r.Erorr404(app)

	app.Get("/layout", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("index", fiber.Map{
			"Title": "Maunguli jasa tukang bangunan",
		}, "layouts/main")
	})

	log.Fatal(app.Listen(":8080"))
}
