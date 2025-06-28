// /src/form_contact.go (VERSI BARU - TELEGRAM)
package src

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func sendMessageHandler(c *fiber.Ctx) error {
	// Ambil data dari form
	name := c.FormValue("name")
	email := c.FormValue("email")
	subject := c.FormValue("subject")
	message := c.FormValue("message")

	// Ambil konfigurasi Telegram dari environment variables
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	chatID := os.Getenv("TELEGRAM_CHAT_ID")

	if botToken == "" || chatID == "" {
		fmt.Println("Error: Variabel Telegram belum diset")
		return c.Redirect("/contact?error=true")
	}

	// Buat pesan yang akan dikirim ke Telegram
	// Kita menggunakan format mirip Markdown agar teksnya tebal
	var sb strings.Builder
	sb.WriteString("🔔 *Pesan Baru dari Website Maunguli* 🔔\n\n")
	sb.WriteString(fmt.Sprintf("*Nama:* %s\n", name))
	sb.WriteString(fmt.Sprintf("*Email:* %s\n", email))
	sb.WriteString(fmt.Sprintf("*Subjek:* %s\n\n", subject))
	sb.WriteString(fmt.Sprintf("*Pesan:*\n%s", message))

	pesanTerkirim := sb.String()

	// Kirim pesan ke API Telegram
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

	_, err := http.PostForm(apiURL, url.Values{
		"chat_id":    {chatID},
		"text":       {pesanTerkirim},
		"parse_mode": {"Markdown"}, // Mengizinkan format tebal, miring, dll.
	})

	if err != nil {
		fmt.Println("Gagal mengirim pesan:", err)
		// Kirim status error dalam format JSON
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal mengirim pesan.",
		})
	}

	fmt.Println("Pesan berhasil dikirim!")
	// Kirim status sukses dalam format JSON
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Pesan berhasil terkirim!",
	})
}

// Pastikan fungsi pendaftaran rute ini dipanggil di main.go
func FormContact(app *fiber.App) {
	app.Post("/send-message", sendMessageHandler)
}
