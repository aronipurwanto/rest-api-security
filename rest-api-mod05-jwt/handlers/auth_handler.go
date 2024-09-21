package handlers

import (
	"github.com/gofiber/fiber/v2"
	"rest-api-mod05-jwt/services"
)

// LoginHandler untuk autentikasi pengguna dan menghasilkan JWT
func LoginHandler(c *fiber.Ctx) error {
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	user, err := services.Authenticate(req.Username, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Generate JWT
	token, err := services.GenerateJWT(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.JSON(fiber.Map{"token": token})
}

// ProtectedHandler hanya dapat diakses dengan JWT yang valid
func ProtectedHandler(c *fiber.Ctx) error {
	claims := c.Locals("claims").(*services.Claims)
	return c.JSON(fiber.Map{"message": "Welcome " + claims.Username + "! You have access."})
}
