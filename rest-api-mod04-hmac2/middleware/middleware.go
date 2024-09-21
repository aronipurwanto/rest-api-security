package middleware

import (
	"github.com/aronipurwanto/rest-api-mod04-hmac2/service"
	"github.com/gofiber/fiber/v2"
	"time"
)

// HMACMiddleware validates the HMAC signature for each request
func HMACMiddleware(c *fiber.Ctx) error {
	// Get the HMAC signature and timestamp from headers
	hmacSignature := c.Get("X-HMAC-Signature")
	timestamp := c.Get("X-Timestamp")

	if hmacSignature == "" || timestamp == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing HMAC signature or timestamp"})
	}

	// Recreate the expected HMAC using the timestamp
	if !service.VerifyHMAC(timestamp, hmacSignature, service.SecretKey) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid HMAC signature"})
	}

	// Check if the timestamp is within 5 minutes to avoid replay attacks
	t, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid timestamp format"})
	}

	if time.Since(t) > 5*time.Minute {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Request expired"})
	}

	// Proceed to the next handler if HMAC is valid
	return c.Next()
}
