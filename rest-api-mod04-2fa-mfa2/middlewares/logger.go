package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

// RequestLogger middleware untuk mencatat setiap request
func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()

		// Tampilkan log request method, path, dan waktu eksekusi
		log.Printf("[%s] %s %s | %v", c.IP(), c.Method(), c.Path(), time.Since(start))

		return err
	}
}
