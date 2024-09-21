package main

import (
	"github.com/aronipurwanto/rest-api-mod04-2fa-mfa/auth"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	// Rute untuk registrasi pengguna
	app.Post("/register", auth.RegisterHandler)

	// Rute untuk login
	app.Post("/login", auth.LoginHandler)

	// Rute yang dilindungi oleh autentikasi 2FA
	app.Get("/secure-data", auth.TwoFAMiddleware, auth.SecureDataHandler)

	// Jalankan server
	log.Fatal(app.Listen(":3000"))
}
