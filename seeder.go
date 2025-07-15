//go:build ignore

package main

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/joho/godotenv"
)

// Mendefinisikan struktur data agar lebih rapi dan aman
type Project struct {
	ProjectTitle string   `firestore:"ProjectTitle"`
	HeroImage    string   `firestore:"HeroImage"`
	Client       string   `firestore:"Client"`
	Location     string   `firestore:"Location"`
	Duration     string   `firestore:"Duration"`
	Services     string   `firestore:"Services"`
	Images       []string `firestore:"Images"`
	Testimonial  string   `firestore:"Testimonial"`
}

type Service struct {
	ServiceName           string   `firestore:"ServiceName"`
	ServiceSubtitle       string   `firestore:"ServiceSubtitle"`
	HeroImage             string   `firestore:"HeroImage"`
	Description           string   `firestore:"Description"`
	Benefits              []string `firestore:"Benefits"`
	RelatedPortfolioSlugs []string `firestore:"RelatedPortfolioSlugs"`
	YouTubeVideoID        string   `firestore:"YouTubeVideoID"`
}

func main() {
	// --- Bagian Koneksi (Sama seperti di main.go) ---
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Println("Peringatan: Tidak bisa memuat file .env")
	}

	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		projectID = "golang-p" // Ganti dengan Project ID Anda
	}

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Gagal terhubung ke Firestore: %v", err)
	}
	defer client.Close()

	log.Println("Berhasil terhubung ke Firestore. Memulai seeding...")

	// --- Data Awal ---
	projectsToSeed := map[string]Project{
		"renovasi-dapur-modern": {
			ProjectTitle: "Renovasi Dapur Modern & Fungsional", HeroImage: "/images/posts/dapur-modern-minimali.png", Client: "Keluarga Santoso", Location: "Jakarta Selatan", Duration: "3 Minggu", Services: "Renovasi Dapur, Pemasangan Granit",
			Images:      []string{"/static/images/portfolio/dapur-1.jpg", "/static/images/portfolio/dapur-2.jpg", "/static/images/portfolio/dapur-3.jpg"},
			Testimonial: "Tim Maunguli mengubah dapur kami yang sempit menjadi luar biasa! Sangat profesional dan hasilnya rapi.",
		},
		"pembangunan-rumah-bsd": {
			ProjectTitle: "Pembangunan Rumah Minimalis di BSD", HeroImage: "/static/images/portfolio/rumah-bsd-hero.jpg", Client: "Keluarga Hermawan", Location: "BSD City, Tangerang Selatan", Duration: "4 Bulan", Services: "Bangun dari Nol, Desain Arsitektur, Interior",
			Images:      []string{"/static/images/portfolio/rumah-bsd-1.jpg", "/static/images/portfolio/rumah-bsd-2.jpg", "/static/images/portfolio/rumah-bsd-3.jpg"},
			Testimonial: "Kami percayakan pembangunan rumah pertama kami pada Maunguli dan hasilnya sangat memuaskan.",
		},
		"fasad-ruko-bintaro": {
			ProjectTitle: "Renovasi Fasad Ruko Komersial", HeroImage: "/static/images/portfolio/ruko-hero.jpg", Client: "Toko Kopi 'Senja'", Location: "Bintaro, Tangerang Selatan", Duration: "1 Bulan", Services: "Renovasi Fasad, Pemasangan ACP, Pengecatan",
			Images:      []string{"/static/images/portfolio/ruko-1.jpg", "/static/images/portfolio/ruko-2.jpg", "/static/images/portfolio/ruko-3.jpg"},
			Testimonial: "Wajah baru ruko kami berhasil menarik lebih banyak pengunjung. Pengerjaan cepat dan hasilnya modern.",
		},
	}

	servicesToSeed := map[string]Service{
		"pasang-granit": {
			ServiceName: "Pasang Granit Lantai & Dinding", ServiceSubtitle: "Keahlian presisi untuk hasil akhir yang mewah.", HeroImage: "images/posts/dapur-modern-minimali.png",
			Description:           "Kami menyediakan jasa pemasangan granit dan marmer untuk lantai, dinding, meja dapur, dan area lainnya.",
			Benefits:              []string{"Hasil Rata & Presisi", "Nat Rapi & Tipis", "Pengerjaan Cepat", "Garansi Pemasangan"},
			RelatedPortfolioSlugs: []string{"renovasi-dapur-modern", "fasad-ruko-bintaro"},
			YouTubeVideoID:        "_IXXYrGlgWE",
		},
		"renovasi-rumah": {
			ServiceName: "Renovasi Total & Parsial", ServiceSubtitle: "Segarkan kembali tampilan properti Anda.", HeroImage: "/static/images/services/renovasi-hero.jpg",
			Description:           "Layanan renovasi lengkap mulai dari perubahan tata ruang, perbaikan struktur, hingga pembaruan interior dan eksterior.",
			Benefits:              []string{"Konsultasi Desain Gratis", "Anggaran Transparan", "Manajemen Proyek Profesional", "Hasil Sesuai Jadwal"},
			RelatedPortfolioSlugs: []string{"renovasi-dapur-modern", "fasad-ruko-bintaro"},
			YouTubeVideoID:        "uOlLl1FquWY",
		},

		"bangun-rumah": {
			ServiceName:           "Bangun Rumah dari Nol",
			ServiceSubtitle:       "Layanan lengkap dari desain hingga serah terima kunci.",
			HeroImage:             "/static/images/portfolio/proyek_struktur_pondasi.webp",
			Description:           "Mewujudkan rumah impian Anda adalah spesialisasi kami. Kami menangani seluruh proses pembangunan, mulai dari perencanaan lahan, konstruksi pondasi, hingga finishing interior dan eksterior, memastikan setiap detail sesuai dengan visi Anda.",
			Benefits:              []string{"Perencanaan Anggaran Detail", "Struktur Kokoh & Aman", "Desain Sesuai Keinginan", "Proses Tepat Waktu"},
			RelatedPortfolioSlugs: []string{"pembangunan-rumah-bsd"},
			YouTubeVideoID:        "yHe4wMykPSE",
		},
	}

	// --- Proses Seeding ---
	// Seeding koleksi 'projects'
	for slug, projectData := range projectsToSeed {
		_, err := client.Collection("projects").Doc(slug).Set(ctx, projectData)
		if err != nil {
			log.Fatalf("Gagal menambahkan proyek %s: %v", slug, err)
		}
		log.Printf("Berhasil menambahkan proyek: %s", slug)
	}

	// Seeding koleksi 'services'
	for slug, serviceData := range servicesToSeed {
		_, err := client.Collection("services").Doc(slug).Set(ctx, serviceData)
		if err != nil {
			log.Fatalf("Gagal menambahkan layanan %s: %v", slug, err)
		}
		log.Printf("Berhasil menambahkan layanan: %s", slug)
	}

	log.Println("Seeding Selesai!")
}
