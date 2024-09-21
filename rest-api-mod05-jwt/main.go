package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
	"log"
	"rest-api-mod05-jwt/config"
	"rest-api-mod05-jwt/handlers"
	"rest-api-mod05-jwt/middleware"
	"rest-api-mod05-jwt/models"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Connect to MySQL database
	config.ConnectDB()

	// Migrate the schema
	config.DB.AutoMigrate(&models.User{}, &models.Role{})

	// Seed data for roles and users
	models.SeedData(config.DB)

	// Initialize Fiber app
	app := fiber.New()

	// Middleware for logging requests
	app.Use(logger.New())

	// Route for login
	app.Post("/login", handlers.LoginHandler)

	// Protected route (requires JWT)
	app.Get("/protected", middleware.JWTMiddleware, handlers.ProtectedHandler)

	// Start the server
	port := fmt.Sprintf(":%d", viper.GetInt("app.port"))
	log.Fatal(app.Listen(port))
}
