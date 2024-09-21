package middleware

import (
	"github.com/gofiber/fiber/v2"
	"rest-api-mod05-auth/models"
	"rest-api-mod05-auth/services"
)

// AuthMiddleware untuk autentikasi user
func AuthMiddleware(c *fiber.Ctx) error {
	username := c.Get("X-Username")
	user, authenticated := services.Authenticate(username)
	if !authenticated {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}
	// Simpan user ke dalam context agar bisa diakses di handler
	c.Locals("user", user)
	return c.Next()
}

// RoleMiddleware untuk otorisasi berdasarkan role
func RoleMiddleware(role string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(models.User)
		if user.Role.Name != role {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Forbidden: Insufficient role"})
		}
		return c.Next()
	}
}

// AuthorityMiddleware untuk otorisasi berdasarkan authorities
func AuthorityMiddleware(authority string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(models.User)
		if !services.HasAuthority(user, authority) {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Forbidden: Insufficient authority"})
		}
		return c.Next()
	}
}
