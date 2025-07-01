// File: /src/routes/services_detail.go
// (VERSI BERSIH - Mengambil data layanan & portofolio terkait dari Firestore)

package routes

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/firestore" // Pastikan import ini ada
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/iterator"
)

// Pastikan Anda memanggil ServicesDetail(app) di main.go
func ServicesDetail(app *fiber.App) {
	// Rute untuk halaman utama galeri semua layanan
	app.Get("/services", servicesIndexHandler)
	// Rute untuk halaman detail setiap layanan
	app.Get("/layanan/:slug", serviceDetailHandler)
}

// Handler untuk halaman utama galeri layanan (/services)
func servicesIndexHandler(c *fiber.Ctx) error {
	ctx := context.Background()
	services := []fiber.Map{}

	// 1. Mengambil semua dokumen dari koleksi 'services'
	iter := firestoreClient.Collection("services").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Gagal mengambil data layanan: %v", err)
			return c.Status(500).SendString("Internal Server Error")
		}

		serviceData := doc.Data()
		serviceData["Slug"] = doc.Ref.ID // Tambahkan slug dari ID dokumen
		services = append(services, serviceData)
	}

	// 2. Render halaman 'services.html' dengan data semua layanan
	return c.Render("services", fiber.Map{
		"Title":    "Layanan Kami - Maunguli",
		"Year":     time.Now().Year(),
		"PageCSS":  "/static/css/page/services.css",
		"Services": services,
	})
}

type ServiceData struct {
	ServiceName           string   `firestore:"ServiceName"`
	ServiceSubtitle       string   `firestore:"ServiceSubtitle"`
	HeroImage             string   `firestore:"HeroImage"`
	Description           string   `firestore:"Description"`
	Benefits              []string `firestore:"Benefits"`
	RelatedPortfolioSlugs []string `firestore:"RelatedPortfolioSlugs"`
	YouTubeVideoID        string   `firestore:"YouTubeVideoID"`
}

// GANTI SELURUH FUNGSI serviceDetailHandler DENGAN INI
func serviceDetailHandler(c *fiber.Ctx) error {
	ctx := context.Background()
	slug := c.Params("slug")

	serviceDoc, err := firestoreClient.Collection("services").Doc(slug).Get(ctx)
	if err != nil {
		log.Printf("Gagal menemukan layanan '%s': %v", slug, err)
		return c.Status(404).Render("404", fiber.Map{"Title": "Layanan Tidak Ditemukan"})
	}

	var service ServiceData
	if err := serviceDoc.DataTo(&service); err != nil {
		log.Printf("Gagal mem-parsing data layanan: %v", err)
		return c.Status(500).SendString("Internal Server Error")
	}

	log.Printf("Data layanan yang berhasil di-load: %+v\n", service)

	relatedProjects := []fiber.Map{}

	// Logika menjadi lebih sederhana karena kita sudah tahu tipenya adalah []string
	if len(service.RelatedPortfolioSlugs) > 0 {
		var docRefs []*firestore.DocumentRef
		for _, projectSlug := range service.RelatedPortfolioSlugs {
			if projectSlug != "" {
				docRef := firestoreClient.Collection("projects").Doc(projectSlug)
				docRefs = append(docRefs, docRef)
			}
		}

		if len(docRefs) > 0 {
			docs, _ := firestoreClient.GetAll(ctx, docRefs)
			for _, doc := range docs {
				if doc.Exists() {
					projectData := doc.Data()
					projectData["Slug"] = doc.Ref.ID
					relatedProjects = append(relatedProjects, projectData)
				}
			}
		}
	}

	// Siapkan data final untuk template
	return c.Render("service-detail", fiber.Map{
		"Title":           service.ServiceName + " - Maunguli",
		"PageCSS":         "/static/css/page/service-detail.css",
		"Year":            time.Now().Year(),
		"ServiceName":     service.ServiceName,
		"ServiceSubtitle": service.ServiceSubtitle,
		"HeroImage":       service.HeroImage,
		"Description":     service.Description,
		"Benefits":        service.Benefits,
		"RelatedProjects": relatedProjects,
		"YoutubeVideoID":  service.YouTubeVideoID,
	})
}
