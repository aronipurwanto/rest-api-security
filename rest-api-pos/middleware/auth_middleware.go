package middleware

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"rest-api-pos/config"
	"rest-api-pos/controller"
	"rest-api-pos/response"
	"strings"
	"time"
)

func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse(401, "Missing Authorization header"))
	}

	// Menghapus "Bearer " dari token
	token := strings.Replace(authHeader, "Bearer ", "", 1)

	// Verifikasi token dengan Keycloak
	idToken, err := config.OIDCVerifier.Verify(context.Background(), token)
	if err != nil {
		// Jika token kadaluwarsa, lakukan refresh token
		if err.Error() == "oidc: token is expired" {
			refreshToken := c.Cookies("refresh_token") // Mengambil refresh token dari cookie (atau simpan di storage lain)

			// Panggil fungsi refreshToken langsung dari controller untuk mendapatkan token baru
			tokenResponse, refreshErr := controller.RefreshTokenFromKeycloak(refreshToken)
			if refreshErr != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse(401, "Token expired and failed to refresh"))
			}

			// Set token baru di header dan cookie, lalu lanjutkan request
			c.Set("Authorization", "Bearer "+tokenResponse.AccessToken)
			c.Cookie(&fiber.Cookie{
				Name:    "refresh_token",
				Value:   tokenResponse.RefreshToken,
				Expires: time.Now().Add(time.Duration(tokenResponse.ExpiresIn) * time.Second),
			})

			// Verifikasi ulang token yang baru
			idToken, err = config.OIDCVerifier.Verify(context.Background(), tokenResponse.AccessToken)
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse(401, "Failed to verify refreshed token"))
			}
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse(401, "Invalid token"))
		}
	}

	// Menyimpan informasi pengguna di context
	var claims map[string]interface{}
	if err := idToken.Claims(&claims); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(500, "Failed to parse token claims"))
	}

	// Simpan claims (informasi user) ke context Fiber
	c.Locals("userClaims", claims)
	return c.Next()
}
