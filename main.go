package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Create a new engine
	engine := html.New("./views", ".html")

	// Or from an embedded system
	// See github.com/gofiber/embed for examples
	// engine := html.NewFileSystem(http.Dir("./views"), ".html")

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./static")

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/main")
	})

	// Custom 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).Render("404", fiber.Map{
			"Title": "404 Not Found",
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
