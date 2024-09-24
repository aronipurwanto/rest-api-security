package middleware

import (
	"encoding/base64"
	"github.com/gofiber/fiber/v2"
	"log"
	"rest-api-mod04-basic-auth/models"
	"strings"
)

// BasicAuth middleware to handle basic authentication
func BasicAuth(c *fiber.Ctx) error {
	log.Println("BasicAuth middleware triggered.")

	// Extract Authorization header
	auth := c.Get("Authorization")
	if auth == "" {
		log.Println("No Authorization header provided.")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    fiber.StatusUnauthorized,
			"message": "Unauthorized - Missing Authorization Header",
		})
	}

	// Check if the authorization header is Basic Auth
	if !strings.HasPrefix(auth, "Basic ") {
		log.Println("Invalid authorization header format.")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    fiber.StatusUnauthorized,
			"message": "Unauthorized - Invalid Authorization Format",
		})
	}

	// Decode the base64 encoded credentials
	decoded, err := base64.StdEncoding.DecodeString(auth[6:])
	if err != nil {
		log.Println("Failed to decode credentials:", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    fiber.StatusUnauthorized,
			"message": "Unauthorized - Failed to Decode Credentials",
		})
	}

	// Split the decoded credentials into username and password
	credentials := strings.Split(string(decoded), ":")
	if len(credentials) != 2 {
		log.Println("Invalid credentials format.")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    fiber.StatusUnauthorized,
			"message": "Unauthorized - Invalid Credentials Format",
		})
	}
	username := credentials[0]
	password := credentials[1]

	log.Printf("Attempting authentication for user: %s", username)

	// Validate user credentials
	if models.ValidateUser(username, password) {
		log.Println("User authentication successful.")
		c.Locals("user", username) // Set authenticated user
		return c.Next()            // Continue to the next handler
	} else {
		log.Println("User authentication failed.")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    fiber.StatusUnauthorized,
			"message": "Unauthorized - Invalid Username or Password",
		})
	}
}
