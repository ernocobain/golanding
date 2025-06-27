package main

import (
	"log"
	"os"

	m "github/dhikrama/go/src"
	// md "github/dhikrama/go/src/midelwares"
	r "github/dhikrama/go/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Create a new engine
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		ViewsLayout: "layouts/main",
		// Prefork:              true,
		Views: engine,
		// CompressedFileSuffix: ".gz",
	})
	app.Use(compress.New(
		compress.Config{
			Level: compress.LevelBestSpeed,
		},
	))

	app.Use(helmet.New(helmet.Config{
		// Konfigurasi lain tetap sama
		HSTSMaxAge:            300,
		HSTSExcludeSubdomains: true,
		XSSProtection:         "0",
		XFrameOptions:         "SAMEORIGIN",

		ContentSecurityPolicy: "default-src 'self'; " +
			"script-src 'self' 'unsafe-inline'; " +
			"style-src 'self' 'unsafe-inline' fonts.googleapis.com; " +
			"font-src 'self' fonts.gstatic.com; " +
			"img-src 'self' data: i.pravatar.cc images.unsplash.com www.transparenttextures.com https://placehold.co;",
		CrossOriginOpenerPolicy:   "unsafe-none",
		CrossOriginEmbedderPolicy: "unsafe-none",
	}))

	app.Static("/static", "./static")

	r.Index(app)
	m.FormContact(app)
	r.Contact(app)
	r.Privacy(app)
	r.About(app)
	r.Services(app)
	r.Erorr404(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen(":" + port))
}
