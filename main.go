// File: main.go (PENDEKATAN TERAKHIR)
package main

import (
	"log"
	"os"
	"strings"

	m "github/dhikrama/go/src"
	r "github/dhikrama/go/src/routes"

	"html/template"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/template/html/v2"
)

func buildCSP(cdnDomains []string) string {
	// Tambahkan domain CDN ke setiap bagian yang diperlukan
	joinedCDN := strings.Join(cdnDomains, " ")

	return strings.Join([]string{
		"default-src 'self';",
		"script-src 'self' 'unsafe-inline' " + joinedCDN + ";",
		"style-src 'self' 'unsafe-inline' fonts.googleapis.com " + joinedCDN + ";",
		"font-src 'self' fonts.gstatic.com;",
		"img-src 'self' data: i.pravatar.cc images.unsplash.com www.transparenttextures.com https://placehold.co " + joinedCDN + ";",
		"frame-src 'self' www.youtube.com;",
	}, " ")
}

func main() {
	r.InitFirestore()

	engine := html.New("./views", ".html")
	engine.Reload(true)

	engine.AddFunc("safeHTML", func(s string) template.HTML {
		return template.HTML(s)
	})

	app := fiber.New(fiber.Config{
		ViewsLayout: "layouts/main",
		Views:       engine,
	})

	app.Use(cors.New(cors.Config{
		// Izinkan domain spesifik. Pisahkan dengan koma.
		// Ini mengizinkan blog Anda dan server pengembangan lokal.
		AllowOrigins: "https://blog.maunguli.com",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	cdnList := []string{
		"https://cdn-maunguli.web.app",
		"https://cdn.maunguli.com",
	}
	csp := buildCSP(cdnList)

	app.Use(helmet.New(helmet.Config{
		XSSProtection:             "0",
		XFrameOptions:             "SAMEORIGIN",
		HSTSMaxAge:                300,
		HSTSExcludeSubdomains:     true,
		ContentSecurityPolicy:     csp, // <-- Menggunakan variabel
		CrossOriginOpenerPolicy:   "unsafe-none",
		CrossOriginEmbedderPolicy: "unsafe-none",
	}))

	// Ganti middleware lama Anda dengan yang ini di main.go
	app.Use(func(c *fiber.Ctx) error {
		hostname := c.Hostname()
		accept := c.Get("Accept")

		// Jika request ke URL Cloud Run langsung → noindex
		if strings.Contains(hostname, ".run.app") {
			c.Set("X-Robots-Tag", "noindex, nofollow")
		}

		// Jika hostname adalah cdn.maunguli.com dan request-nya bukan image
		if strings.Contains(hostname, "cdn.maunguli.com") && !strings.HasPrefix(accept, "image/") {
			c.Set("X-Robots-Tag", "noindex, nofollow")
		}

		return c.Next()
	})

	app.Static("/", "./public")

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
