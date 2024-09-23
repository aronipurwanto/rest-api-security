package middlewares

import "github.com/gofiber/fiber/v2"

// CORSMiddleware adds CORS headers to allow cross-origin requests
func CORSMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Set CORS headers
		c.Set("Access-Control-Allow-Origin", "*") // You can restrict this to specific origins
		c.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")
		c.Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight request (OPTIONS method)
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusOK)
		}

		return c.Next()
	}
}
