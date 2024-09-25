package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"rest-api-pos/config"
	"rest-api-pos/response"
)

// TokenResponse represents the response from Keycloak
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

// Login handles the login to Keycloak and returns access token and refresh token
func Login(c *fiber.Ctx) error {
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var loginRequest LoginRequest
	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(400, "Invalid request"))
	}

	tokenResponse, err := getTokenFromKeycloak(loginRequest.Username, loginRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(500, "Failed to login"))
	}

	return c.JSON(response.SuccessResponse(tokenResponse, "Login successful"))
}

// getTokenFromKeycloak fetches token from Keycloak using username and password
func getTokenFromKeycloak(username, password string) (*TokenResponse, error) {
	data := map[string]string{
		"client_id":     config.AppConfig.Keycloak.ClientID,
		"client_secret": config.AppConfig.Keycloak.ClientSecret,
		"grant_type":    "password",
		"username":      username,
		"password":      password,
	}

	reqBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(config.AppConfig.Keycloak.URL+config.AppConfig.Keycloak.TokenEndpoint, "application/x-www-form-urlencoded", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	var tokenResponse TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return nil, err
	}

	return &tokenResponse, nil
}

// RefreshToken refreshes the access token using refresh token
func RefreshToken(c *fiber.Ctx) error {
	type RefreshRequest struct {
		RefreshToken string `json:"refresh_token"`
	}

	var refreshRequest RefreshRequest
	if err := c.BodyParser(&refreshRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(400, "Invalid request"))
	}

	tokenResponse, err := RefreshTokenFromKeycloak(refreshRequest.RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(500, "Failed to refresh token"))
	}

	return c.JSON(response.SuccessResponse(tokenResponse, "Token refreshed successfully"))
}

// RefreshTokenFromKeycloak refreshes the access token using refresh token
func RefreshTokenFromKeycloak(refreshToken string) (*TokenResponse, error) {
	data := map[string]string{
		"client_id":     config.AppConfig.Keycloak.ClientID,
		"client_secret": config.AppConfig.Keycloak.ClientSecret,
		"grant_type":    "refresh_token",
		"refresh_token": refreshToken,
	}

	reqBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(config.AppConfig.Keycloak.URL+config.AppConfig.Keycloak.TokenEndpoint, "application/x-www-form-urlencoded", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	var tokenResponse TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return nil, err
	}

	return &tokenResponse, nil
}

// GetProfile handles the request to get the user profile based on the logged-in user
func GetProfile(c *fiber.Ctx) error {
	// Ambil token dari header Authorization
	accessToken := c.Get("Authorization")

	// Pastikan token diawali dengan "Bearer "
	if len(accessToken) > 7 && accessToken[:7] == "Bearer " {
		accessToken = accessToken[7:] // Hapus "Bearer " dari token
	}

	// Panggil Keycloak untuk mendapatkan user profile menggunakan access token
	profile, err := getUserProfileFromKeycloak(accessToken)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(500, "Failed to fetch user profile"))
	}

	// Kembalikan profil pengguna
	return c.JSON(response.SuccessResponse(profile, "User profile fetched successfully"))
}

// getUserProfileFromKeycloak fetches the user profile from Keycloak using the access token
func getUserProfileFromKeycloak(accessToken string) (map[string]interface{}, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", config.AppConfig.Keycloak.URL+"/realms/"+config.AppConfig.Keycloak.Realm+"/protocol/openid-connect/userinfo", nil)
	if err != nil {
		return nil, err
	}

	// Tambahkan header Authorization dengan token
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Dekode hasil dari response menjadi map
	var profile map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return nil, err
	}

	return profile, nil
}
