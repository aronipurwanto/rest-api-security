package middleware

import (
	"github.com/gofiber/fiber/v2"
	"rest-api-security-mod06-logging/utils"
	"time"
)

// LoggingMiddleware adalah middleware untuk mencatat permintaan dan respons
func LoggingMiddleware(c *fiber.Ctx) error {
	start := time.Now()

	// Log request
	utils.Log.WithFields(utils.Fields{
		"method":   c.Method(),
		"endpoint": c.OriginalURL(),
		"ip":       c.IP(),
		"headers":  c.GetReqHeaders(),
	}).Info("Permintaan API diterima")

	// Proses request
	err := c.Next()

	// Log response
	duration := time.Since(start)
	utils.Log.WithFields(utils.Fields{
		"status":   c.Response().StatusCode(),
		"duration": duration.Seconds() * 1000,
		"size":     len(c.Response().Body()),
	}).Info("Respons API")

	return err
}
