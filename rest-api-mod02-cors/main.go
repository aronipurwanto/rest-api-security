package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"rest-api-mod02-cors/controllers"
	"rest-api-mod02-cors/middlewares"
	"rest-api-mod02-cors/services"
)

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Initialize session service
	sessionService := services.NewSessionService()

	// Initialize user controller
	userController := controllers.UserController{SessionService: sessionService}

	// Apply CORS middleware
	app.Use(middlewares.CORSMiddleware())

	// Define routes for login, logout, and session check
	app.Post("/login", userController.LoginHandler)
	app.Post("/logout", userController.LogoutHandler)
	app.Get("/session", userController.SessionCheckHandler)

	// Start the server
	fmt.Println("Server running on port 3000...")
	log.Fatal(app.Listen(":3000"))
}
