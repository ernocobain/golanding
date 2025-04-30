package main

import (
	"log"
	"time"

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
		HSTSMaxAge:            300,
		HSTSExcludeSubdomains: true,
		XSSProtection:         "0",
		XFrameOptions:         "ALLOW-FROM https://maunguli.com",
	}))
	app.Static("/", "./static")

	r.Index(app)

	m.FormContact(app)

	// Custom 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).Render("404", fiber.Map{
			"Title": "404 Not Found",
			"Year":  time.Now().Year(),
		}, "layouts/main") // Render the 404 template
	})

	app.Get("/layout", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("index", fiber.Map{
			"Title": "Maunguli jasa tukang bangunan",
		}, "layouts/main")
	})

	app.Get("/layouts-nested", func(c *fiber.Ctx) error {
		// Render index within layouts/nested/main within layouts/nested/base
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/nested/main", "layouts/nested/base")
	})

	log.Fatal(app.Listen(":8080"))
}
