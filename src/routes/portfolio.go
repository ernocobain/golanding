// /src/routes/portfolio.go
package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// Pastikan Anda memanggil Portfolio(app) di main.go
func Portfolio(app *fiber.App) {
	app.Get("/portfolio", portfolioIndexHandler)
	// Rute dinamis dengan parameter 'slug'
	app.Get("/portfolio/:slug", portfolioDetailHandler)
}

// Data proyek (untuk sementara, idealnya ini dari database)
var projects = map[string]fiber.Map{
	"renovasi-dapur-modern": {
		"Title":        "Studi Kasus: Renovasi Dapur Modern",
		"ProjectTitle": "Renovasi Dapur Modern & Fungsional",
		"HeroImage":    "/static/images/portfolio/dapur-hero.jpg",
		"Client":       "Keluarga Santoso",
		"Location":     "Jakarta Selatan",
		"Duration":     "3 Minggu",
		"Services":     "Renovasi Dapur, Pemasangan Granit",
		"Images": []string{
			"/static/images/portfolio/dapur-1.jpg",
			"/static/images/portfolio/dapur-2.jpg",
			"/static/images/portfolio/dapur-3.jpg",
		},
		"Testimonial": "Tim Maunguli mengubah dapur kami yang sempit menjadi luar biasa! Sangat profesional dan hasilnya rapi.",
	},
	// Tambahkan data proyek lain di sini dengan slug yang unik
	// /src/routes/portfolio.go

	// --- TAMBAHKAN DUA PROYEK BARU DI BAWAH INI ---

	// Contoh Proyek 2: Pembangunan Rumah
	"pembangunan-rumah-bsd": {
		"ProjectTitle": "Pembangunan Rumah Minimalis di BSD",
		"HeroImage":    "/static/images/portfolio/rumah-bsd-hero.jpg",
		"Client":       "Keluarga Hermawan",
		"Location":     "BSD City, Tangerang Selatan",
		"Duration":     "4 Bulan",
		"Services":     "Bangun dari Nol, Desain Arsitektur, Interior",
		"Images": []string{
			"/static/images/portfolio/rumah-bsd-1.jpg",
			"/static/images/portfolio/rumah-bsd-2.jpg",
			"/static/images/portfolio/rumah-bsd-3.jpg",
		},
		"Testimonial": "Kami percayakan pembangunan rumah pertama kami pada Maunguli dan hasilnya sangat memuaskan. Prosesnya transparan dan komunikasinya sangat baik dari awal hingga akhir.",
	},

	// Contoh Proyek 3: Renovasi Komersial
	"fasad-ruko-bintaro": {
		"ProjectTitle": "Renovasi Fasad Ruko Komersial",
		"HeroImage":    "/static/images/portfolio/ruko-hero.jpg",
		"Client":       "Toko Kopi 'Senja'",
		"Location":     "Bintaro, Tangerang Selatan",
		"Duration":     "1 Bulan",
		"Services":     "Renovasi Fasad, Pemasangan ACP, Pengecatan",
		"Images": []string{
			"/static/images/portfolio/ruko-1.jpg",
			"/static/images/portfolio/ruko-2.jpg",
			"/static/images/portfolio/ruko-3.jpg",
		},
		"Testimonial": "Wajah baru ruko kami berhasil menarik lebih banyak pengunjung. Pengerjaan dari tim Maunguli cepat dan hasilnya sangat modern. Terima kasih!",
	},
}

// Handler BARU untuk halaman galeri (/portfolio)
func portfolioIndexHandler(c *fiber.Ctx) error {
	return c.Render("portfolio-index", fiber.Map{
		"Title":    "Portofolio Proyek - Maunguli",
		"Year":     time.Now().Year(),
		"PageCSS":  "/static/css/page/portfolio.css",
		"Projects": projects, // <-- Mengirim semua data proyek ke template
	})
}

func portfolioDetailHandler(c *fiber.Ctx) error {
	slug := c.Params("slug")
	project, ok := projects[slug]

	// Jika proyek tidak ditemukan, arahkan ke halaman 404
	if !ok {
		return c.Status(404).Render("404", fiber.Map{"Title": "Proyek Tidak Ditemukan"})
	}

	// Menambahkan data standar ke data proyek
	project["Year"] = time.Now().Year()
	project["PageCSS"] = "/static/css/page/portfolio-detail.css"
	project["Title"] = project["ProjectTitle"].(string) + " - Maunguli"

	// Render template portfolio-detail.html dengan data proyek yang sesuai
	return c.Render("portfolio-detail", project)
}
