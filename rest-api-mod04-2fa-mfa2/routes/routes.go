package routes

import (
	"github.com/gofiber/fiber/v2"
	"rest-api-mod04-2fa-mfa2/handlers"
	"rest-api-mod04-2fa-mfa2/middlewares"
)

// SetupRoutes untuk mengatur route di Fiber
func SetupRoutes(app *fiber.App) {
	// Route untuk login (step 1 - kirim OTP)
	app.Post("/login", handlers.Login)

	// Route untuk verifikasi OTP (step 2 - verifikasi 2FA)
	app.Post("/verify-otp", handlers.VerifyOTP)

	// Group route yang membutuhkan autentikasi JWT
	authRequired := app.Group("/", middlewares.JWTMiddleware())

	// Endpoint terproteksi
	authRequired.Get("/profile", handlers.Profile)
	authRequired.Get("/settings", handlers.Settings)
	authRequired.Get("/dashboard", handlers.Dashboard)

	// Route baru untuk mengekstrak token
	app.Get("/extract-token", handlers.ExtractToken)
}
