package midelwares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CssPurge(c *fiber.Ctx) error {
	// Handle CSS files
	if strings.HasSuffix(c.Path(), ".css") {
		c.Set("Content-Type", "text/css")
		// Set cache headers
		c.Set("Cache-Control", "public, max-age=31536000, immutable")
	}

	return c.Next()
}
