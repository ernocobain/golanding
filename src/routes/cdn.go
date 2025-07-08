package routes

import (
	"bytes"
	"image"
	"net/http"
	"strconv"
	"strings"

	webp "github.com/chai2010/webp"
	"github.com/disintegration/imaging"
	"github.com/gofiber/fiber/v2"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func CdnHandler(app *fiber.App) {
	app.Get("/cdn", func(c *fiber.Ctx) error {
		src := c.Query("src")
		widthStr := c.Query("w")
		heightStr := c.Query("h")
		format := c.Query("format", "webp")

		if src == "" {
			return c.Status(400).SendString("Missing `src` parameter")
		}

		// Tambahkan domain default jika src relatif
		if !strings.HasPrefix(src, "http") {
			src = "https://cdn.maunguli.com" + src
		}

		// Ambil gambar dari URL
		resp, err := http.Get(src)
		if err != nil || resp.StatusCode != 200 {
			return c.Status(404).SendString("Image not found")
		}
		defer resp.Body.Close()

		// Decode gambar
		img, _, err := image.Decode(resp.Body)
		if err != nil {
			return c.Status(500).SendString("Decode failed")
		}

		// Ubah ukuran jika diminta
		width, _ := strconv.Atoi(widthStr)
		height, _ := strconv.Atoi(heightStr)
		if width > 0 || height > 0 {
			img = imaging.Resize(img, width, height, imaging.Lanczos)
		}

		// Encode ke format yang diminta
		buf := new(bytes.Buffer)
		switch format {
		case "jpeg":
			c.Type("jpeg")
			err = imaging.Encode(buf, img, imaging.JPEG)
		case "png":
			c.Type("png")
			err = imaging.Encode(buf, img, imaging.PNG)
		default: // webp
			c.Type("webp")
			err = webp.Encode(buf, img, &webp.Options{Quality: 85})
		}
		if err != nil {
			return c.Status(500).SendString("Encode failed")
		}

		c.Set("Cache-Control", "public, max-age=31536000, immutable")
		return c.Send(buf.Bytes())
	})
}
