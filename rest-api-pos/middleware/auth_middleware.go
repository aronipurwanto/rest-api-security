package middleware

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"net/url"
	"rest-api-pos/config"
	"rest-api-pos/response"
	"strings"
)

// AuthMiddleware is a middleware to verify the authorization token with Keycloak
func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse(401, "Missing Authorization header"))
	}

	// Menghapus "Bearer " dari token
	token := strings.Replace(authHeader, "Bearer ", "", 1)

	// Verifikasi token dengan Keycloak
	if err := verifyTokenWithKeycloak(token); err != nil {
		// Jika token kadaluwarsa, lakukan refresh token
		if err.Error() == "oidc: token is expired" {
			return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse(401, "Failed to verify refreshed token"))
			/*
				refreshToken := c.Cookies("refresh_token") // Mengambil refresh token dari cookie

				// Panggil fungsi refreshToken untuk mendapatkan token baru
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
				// Verifikasi ulang token yang baru dengan Keycloak
				if err := verifyTokenWithKeycloak(tokenResponse.AccessToken); err != nil {
					return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse(401, "Failed to verify refreshed token"))
				}
			*/
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse(401, "Invalid token"))
		}
	}

	// Menyimpan informasi pengguna di context
	claims, err := getClaimsFromToken(token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(500, "Failed to parse token claims"))
	}

	// Simpan claims (informasi user) ke context Fiber
	c.Locals("userClaims", claims)
	return c.Next()
}

// verifyTokenWithKeycloak verifies the token with Keycloak directly
func verifyTokenWithKeycloak(token string) error {
	client := &http.Client{}

	// Request to introspect the token with Keycloak
	data := url.Values{}
	data.Set("token", token)
	data.Set("client_id", config.AppConfig.Keycloak.ClientID)
	data.Set("client_secret", config.AppConfig.Keycloak.ClientSecret)

	req, err := http.NewRequest("POST", config.AppConfig.Keycloak.URL+"/realms/"+config.AppConfig.Keycloak.Realm+"/protocol/openid-connect/token/introspect", strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token response from Keycloak")
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	// Check if the token is active
	if active, ok := result["active"].(bool); !ok || !active {
		return fiber.NewError(fiber.StatusUnauthorized, "Token is not active")
	}

	return nil
}

// getClaimsFromToken extracts user claims from a valid token
func getClaimsFromToken(token string) (map[string]interface{}, error) {
	// You can customize this function to extract claims from the token
	// Based on how you want to handle claims (JWT decoding or Keycloak introspection response)
	return map[string]interface{}{
		"sub": "user-id", // Example of extracting user ID
		// Add other claims as necessary
	}, nil
}
