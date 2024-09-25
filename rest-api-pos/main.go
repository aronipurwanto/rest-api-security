package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"rest-api-pos/config"
	"rest-api-pos/middleware"
	"rest-api-pos/router"
)

func main() {
	// Load configuration and initialize Keycloak OIDC verifier
	config.LoadConfig()

	// Initialize Fiber app
	app := fiber.New()

	// Add middleware (example: request logging)
	app.Use(middleware.AuthMiddleware)

	// Setup routes
	router.SetupRoutes(app)

	// Start the server
	port := config.AppConfig.App.Port
	log.Printf("Server is running on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
