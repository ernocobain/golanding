package utils

import (
	"fmt"
	"strings"
)

func GeneratePictureHTML(src string, alt string, widths []int, sizes string) string {
	cdnBaseUrl := "https://cdn.maunguli.com" // Gunakan domain CDN Anda

	webpSet := []string{}

	for _, w := range widths {
		// Buat URL yang memanggil layanan image-resizer (format WebP)
		webpSet = append(webpSet, fmt.Sprintf("%s/cdn?src=%s&w=%d&format=webp %dw", cdnBaseUrl, src, w, w))
	}
	defaultW := widths[len(widths)/2] // Ambil resolusi tengah sebagai fallback default

	return fmt.Sprintf(`
<picture>
  <source type="image/webp" srcset="%s" sizes="%s">
  <img src="%s/cdn?src=%s&w=%d&format=webp" width="%d" alt="%s" loading="lazy" decoding="async">
</picture>
`, strings.Join(webpSet, ", "), sizes, cdnBaseUrl, src, defaultW, defaultW, alt)
}
