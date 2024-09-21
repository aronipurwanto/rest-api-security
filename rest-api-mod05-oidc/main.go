package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
	"log"
	"rest-api-mod05-oidc/config"
	"rest-api-mod05-oidc/handlers"
	"rest-api-mod05-oidc/models"
	"rest-api-mod05-oidc/services"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Connect to MySQL database
	config.ConnectDB()

	// Migrate the schema
	config.DB.AutoMigrate(&models.User{})

	// Seed data for users (optional)
	models.SeedData(config.DB)

	// Initialize OAuth config
	services.InitOAuthConfig()

	// Initialize Fiber app
	app := fiber.New()

	// Middleware for logging requests
	app.Use(logger.New())

	// OAuth login route
	app.Get("/oauth/login", handlers.OAuthLogin)

	// OAuth callback route
	app.Get("/oauth/callback", handlers.OAuthCallback)

	// Start the server
	port := fmt.Sprintf(":%d", viper.GetInt("app.port"))
	log.Fatal(app.Listen(port))
}
