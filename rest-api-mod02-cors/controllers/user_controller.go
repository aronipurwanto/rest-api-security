package controllers

import (
	"github.com/gofiber/fiber/v2"
	"rest-api-mod02-cors/services"
)

// UserController struct to handle user-related routes
type UserController struct {
	SessionService services.SessionService
}

// LoginHandler simulates a login action and sets session
func (uc *UserController) LoginHandler(c *fiber.Ctx) error {
	// Simulate a login and set session
	username := c.FormValue("username")
	if username == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    "400",
			"message": "Username is required",
		})
	}

	// Save user in session
	uc.SessionService.SaveSession(c, "username", username)

	// Set a cookie (for example purposes)
	c.Cookie(&fiber.Cookie{
		Name:  "user_cookie",
		Value: username,
	})

	return c.JSON(fiber.Map{
		"code":    "200",
		"message": "Login successful",
		"data":    fiber.Map{"username": username},
	})
}

// LogoutHandler clears the session and cookie
func (uc *UserController) LogoutHandler(c *fiber.Ctx) error {
	// Clear session
	uc.SessionService.ClearSession(c, "username")

	// Clear cookie
	c.Cookie(&fiber.Cookie{
		Name:   "user_cookie",
		Value:  "",
		MaxAge: -1, // Deletes the cookie
	})

	return c.JSON(fiber.Map{
		"code":    "200",
		"message": "Logout successful",
	})
}

// SessionCheckHandler checks if a session exists
func (uc *UserController) SessionCheckHandler(c *fiber.Ctx) error {
	// Check if session exists
	username := uc.SessionService.GetSession(c, "username")
	if username == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    "401",
			"message": "No active session",
		})
	}

	return c.JSON(fiber.Map{
		"code":    "200",
		"message": "Session active",
		"data":    fiber.Map{"username": username},
	})
}
