package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"log"
)

// Kunci rahasia yang sama digunakan untuk menandatangani token
var jwtSecret = []byte("your_secret_key")

// JWTMiddleware untuk memvalidasi JWT token
func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Mendapatkan token dari header Authorization
		tokenString := c.Get("Authorization")

		if tokenString == "" {
			log.Println("Validasi JWT gagal: Token tidak ditemukan")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Missing or invalid token",
			})
		}

		// Hapus kata "Bearer " dari token jika ada
		tokenString = tokenString[len("Bearer "):]

		// Memparsing token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validasi metode tanda tangan token
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.ErrUnauthorized
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			log.Println("Validasi JWT gagal:", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid token",
			})
		}

		// Token valid, lanjutkan ke handler berikutnya
		log.Println("Validasi JWT berhasil")
		return c.Next()
	}
}
