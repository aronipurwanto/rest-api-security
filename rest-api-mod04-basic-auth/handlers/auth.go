package handlers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"rest-api-mod04-basic-auth/models"
)

// LoginWithBody handles login via request body (username and password)
func LoginWithBody(c *fiber.Ctx) error {
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Invalid request body",
		})
	}

	// Validate username and password
	if models.ValidateUser(req.Username, req.Password) {
		log.Printf("User %s successfully authenticated via body login", req.Username)
		return c.JSON(fiber.Map{
			"code":    200,
			"message": "Authentication successful",
			"data":    fiber.Map{"user": req.Username},
		})
	} else {
		log.Println("User authentication failed for user:", req.Username)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    fiber.StatusUnauthorized,
			"message": "Unauthorized - Invalid Username or Password",
		})
	}
}

// LoginHandler handles the login logic after successful Basic Authentication
func LoginHandler(c *fiber.Ctx) error {
	log.Println("LoginHandler reached, user authenticated.")

	// Retrieve authenticated user from the context
	user := c.Locals("user").(string)

	// Log the authenticated user
	log.Printf("User %s successfully authenticated", user)

	// Response after successful authentication
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "Authentication successful",
		"data":    fiber.Map{"user": user},
	})
}

// Profile is a protected endpoint for Basic Auth users
func Profile(c *fiber.Ctx) error {
	user := c.Locals("user").(string)
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "Profile information retrieved successfully",
		"data":    fiber.Map{"user": user, "profile": "Profile data for " + user},
	})
}

// Settings is a protected endpoint for Basic Auth users
func Settings(c *fiber.Ctx) error {
	user := c.Locals("user").(string)
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "Settings retrieved successfully",
		"data":    fiber.Map{"user": user, "settings": "Settings data for " + user},
	})
}

// Dashboard is a protected endpoint for Basic Auth users
func Dashboard(c *fiber.Ctx) error {
	user := c.Locals("user").(string)
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "Dashboard data retrieved successfully",
		"data":    fiber.Map{"user": user, "dashboard": "Dashboard data for " + user},
	})
}
