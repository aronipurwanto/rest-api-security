package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"rest-api-mod05-rbac/models"
	"rest-api-mod05-rbac/services"
)

// AuthMiddleware untuk autentikasi user
func AuthMiddleware(c *fiber.Ctx) error {
	username := c.Get("X-Username")
	user, authenticated := services.Authenticate(username)
	if !authenticated {
		logrus.Warn("Unauthorized access attempt by:", username)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}
	// Simpan user ke dalam context agar bisa diakses di handler
	c.Locals("user", user)
	return c.Next()
}

// RoleMiddleware untuk otorisasi berdasarkan role (RBAC)
func RoleMiddleware(role string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(models.User)
		if user.Role.Name != role {
			logrus.Warn("Forbidden access attempt by:", user.Username)
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Forbidden: Insufficient role"})
		}
		return c.Next()
	}
}

// ACLMiddleware untuk otorisasi berdasarkan ACL
func ACLMiddleware(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)
	path := c.Path()
	method := c.Method()

	if !services.HasAccessControl(user, path, method) {
		logrus.Warn("Forbidden access attempt by:", user.Username, "on path:", path)
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Forbidden: Access denied"})
	}
	return c.Next()
}
