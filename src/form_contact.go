package src

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func FormContact(app *fiber.App) {
	// --- Handler untuk Menerima Data Form Kontak ---
	app.Post("/contact", func(c *fiber.Ctx) error {
		// 1. Ambil data dari form menggunakan atribut 'name'
		name := c.FormValue("name")
		email := c.FormValue("email")
		message := c.FormValue("message")

		// 2. (Contoh) Cetak data ke konsol server
		fmt.Printf("Pesan Diterima:\n Nama: %s\n Email: %s\n Pesan: %s\n", name, email, message)

		// 3. (Penting) Lakukan sesuatu dengan data ini!
		//    - Validasi data (pastikan email valid, nama tidak kosong, dll.)
		//    - Kirim email notifikasi
		//    - Simpan ke database
		//    - dll.

		// 4. Kirim respons kembali ke pengguna
		//    Contoh: Redirect ke halaman terima kasih
		//    return c.Redirect("/thank-you")

		//    Contoh: Kirim pesan sukses sederhana
		return c.SendString("Pesan Anda telah diterima!")

		//    Contoh: Render halaman lain dengan pesan sukses
		//    return c.Render("contact-success", fiber.Map{
		//        "Title": "Pesan Terkirim",
		//        "Year": time.Now().Year(),
		//    }, "layouts/main")
	})
	// --- Akhir Handler Form Kontak ---

}
