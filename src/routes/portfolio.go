// File: /src/routes/portfolio.go (VERSI FINAL & LENGKAP)

package routes

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/iterator"
	// Tidak perlu import "cloud.google.com/go/firestore" jika tidak ada tipe data spesifik yang dipakai di sini
)

// --- TAMBAHKAN BARIS INI UNTUK MENDEKLARASIKAN VARIABEL ---
// var firestoreClient *firestore.Client
// CATATAN: Jika Anda sudah membuat file firestore_client.go, baris ini tidak diperlukan.
// Jika belum, Anda bisa menaruhnya di sini atau di file lain dalam package routes.
// Untuk amannya, kita asumsikan ia sudah ada di package routes.

// Pastikan Anda memanggil Portfolio(app) di main.go
func Portfolio(app *fiber.App) {
	// Rute untuk halaman utama galeri
	app.Get("/portfolio", portfolioIndexHandler)

	// Rute untuk halaman detail
	app.Get("/portfolio/:slug", portfolioDetailHandler)
}

// Handler untuk halaman galeri (/portfolio)
func portfolioIndexHandler(c *fiber.Ctx) error {
	ctx := context.Background()
	projects := make(map[string]fiber.Map)

	// Baris ini sekarang akan berfungsi karena 'firestoreClient' sudah dikenal
	iter := firestoreClient.Collection("projects").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Gagal mengambil data proyek: %v", err)
			return c.Status(500).SendString("Internal Server Error")
		}

		projectData := doc.Data()
		projects[doc.Ref.ID] = projectData
	}

	return c.Render("portfolio-index", fiber.Map{
		"Title":    "Portofolio Proyek - Maunguli",
		"Year":     time.Now().Year(),
		"Projects": projects,
	})
}

// Handler untuk halaman detail (/portfolio/:slug)
func portfolioDetailHandler(c *fiber.Ctx) error {
	ctx := context.Background()
	slug := c.Params("slug")

	doc, err := firestoreClient.Collection("projects").Doc(slug).Get(ctx)
	if err != nil {
		log.Printf("Gagal menemukan proyek '%s': %v", slug, err)
		return c.Status(404).Render("404", fiber.Map{"Title": "Proyek Tidak Ditemukan"})
	}

	projectData := doc.Data()
	projectData["Year"] = time.Now().Year()
	projectData["PageCSS"] = "/static/css/page/portfolio-detail.css"
	if title, ok := projectData["ProjectTitle"].(string); ok {
		projectData["Title"] = title + " - Maunguli"
	} else {
		projectData["Title"] = "Detail Proyek - Maunguli"
	}

	return c.Render("portfolio-detail", projectData)
}
