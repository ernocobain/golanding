package routes

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/joho/godotenv"
)

// Definisikan client sebagai variabel level package (global untuk package routes)
var firestoreClient *firestore.Client

// Fungsi ini akan dipanggil sekali saja dari main.go
func InitFirestore() {
	ctx := context.Background()

	if err := godotenv.Load(); err != nil {
		log.Println("Peringatan: Tidak bisa memuat file .env.")
	}

	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		projectID = "golang-p" // Fallback
	}

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Gagal terhubung ke Firestore: %v", err)
	}
	firestoreClient = client
	log.Println("Berhasil terhubung ke Firestore.")
}
