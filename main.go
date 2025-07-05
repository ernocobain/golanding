// File: main.go (PENDEKATAN TERAKHIR)
package main

import (
	"log"
	"os"

	m "github/dhikrama/go/src"
	r "github/dhikrama/go/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/template/html/v2"
)

func main() {
	r.InitFirestore()

	engine := html.New("./views", ".html")
	engine.Reload(true)

	app := fiber.New(fiber.Config{
		ViewsLayout: "layouts/main",
		Views:       engine,
	})

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	// --- PENDEKATAN BARU UNTUK CSP ---
	// Definisikan string CSP di sini
	var csp = "default-src 'self'; " +
		"script-src 'self' 'unsafe-inline'; " +
		"style-src 'self' 'unsafe-inline' fonts.googleapis.com; " +
		"font-src 'self' fonts.gstatic.com; " +
		"img-src 'self' data: i.pravatar.cc images.unsplash.com www.transparenttextures.com https://placehold.co; " +
		"frame-src 'self' www.youtube.com;"

	// Masukkan variabel CSP ke dalam config
	app.Use(helmet.New(helmet.Config{
		XSSProtection:             "0",
		XFrameOptions:             "SAMEORIGIN",
		HSTSMaxAge:                300,
		HSTSExcludeSubdomains:     true,
		ContentSecurityPolicy:     csp, // <-- Menggunakan variabel
		CrossOriginOpenerPolicy:   "unsafe-none",
		CrossOriginEmbedderPolicy: "unsafe-none",
	}))

	app.Static("/static", "./public/static")

	r.Index(app)
	m.FormContact(app)
	r.Contact(app)
	r.Privacy(app)
	r.About(app)
	r.Services(app)
	r.Portfolio(app)
	r.ServicesDetail(app)
	r.Erorr404(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(app.Listen(":" + port))
}
