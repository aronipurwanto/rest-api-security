package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

// IPWhitelistMiddleware is a middleware function that checks if the client IP is whitelisted
func IPWhitelistMiddleware(allowedIPs []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get client IP
		clientIP := c.IP()

		// Check if client IP is in the whitelist
		for _, ip := range allowedIPs {
			// Allow access if IP matches
			if clientIP == ip {
				return c.Next()
			}
		}

		// If IP is not whitelisted, return 403 Forbidden
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"code":    "403",
			"message": "Your IP is not allowed to access this resource",
			"data":    nil,
		})
	}
}
