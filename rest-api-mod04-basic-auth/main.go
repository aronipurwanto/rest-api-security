package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"rest-api-mod04-basic-auth/handlers"
	"rest-api-mod04-basic-auth/middleware"
	"rest-api-mod04-basic-auth/models"
)

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Initialize sample users in memory
	models.InitializeUsers()

	// Logging to indicate server start
	log.Println("Starting server on port 3000...")

	// Routes for Basic Auth endpoints
	app.Post("/login-body", handlers.LoginWithBody)                        // Login with username and password from body
	app.Post("/login-header", middleware.BasicAuth, handlers.LoginHandler) // Login with Basic Auth header

	// Protected sample endpoints
	app.Get("/profile", middleware.BasicAuth, handlers.Profile)
	app.Get("/settings", middleware.BasicAuth, handlers.Settings)
	app.Get("/dashboard", middleware.BasicAuth, handlers.Dashboard)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
