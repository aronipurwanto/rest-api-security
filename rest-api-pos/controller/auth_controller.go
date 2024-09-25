package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"net/http"
	"net/url"
	"rest-api-pos/config"
	"rest-api-pos/response"
	"strings"
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
	data := url.Values{}
	data.Set("client_id", config.AppConfig.Keycloak.ClientID)
	data.Set("client_secret", config.AppConfig.Keycloak.ClientSecret)
	data.Set("grant_type", "password")
	data.Set("username", username)
	data.Set("password", password)

	keycloakUrl := fmt.Sprintf("%s/%s", config.AppConfig.Keycloak.URL, config.AppConfig.Keycloak.TokenEndpoint)
	// Buat request POST
	req, err := http.NewRequest("POST", keycloakUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	// Set header Content-Type ke form-url-encoded
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Baca body respons
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Periksa tipe konten dari respons
	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		// Jika bukan JSON, tampilkan respons untuk debugging
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Unexpected response format from Keycloak: "+string(bodyBytes))
	}

	// Periksa jika status code tidak 200 OK
	if resp.StatusCode != http.StatusOK {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Keycloak responded with an error: "+string(bodyBytes))
	}

	// Decode JSON dari body response
	var tokenResponse TokenResponse
	if err := json.Unmarshal(bodyBytes, &tokenResponse); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to decode Keycloak response: "+string(bodyBytes))
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
	data := url.Values{}
	data.Set("client_id", config.AppConfig.Keycloak.ClientID)
	data.Set("client_secret", config.AppConfig.Keycloak.ClientSecret)
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)

	keycloakUrl := fmt.Sprintf("%s/%s", config.AppConfig.Keycloak.URL, config.AppConfig.Keycloak.TokenEndpoint)
	req, err := http.NewRequest("POST", keycloakUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
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
	accessToken := c.Get("Authorization")

	if len(accessToken) > 7 && accessToken[:7] == "Bearer " {
		accessToken = accessToken[7:] // Remove "Bearer "
	}

	profile, err := getUserProfileFromKeycloak(accessToken)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(500, "Failed to fetch user profile"))
	}

	return c.JSON(response.SuccessResponse(profile, "User profile fetched successfully"))
}

// getUserProfileFromKeycloak fetches the user profile from Keycloak using the access token
func getUserProfileFromKeycloak(accessToken string) (map[string]interface{}, error) {
	client := &http.Client{}
	keycloakUrl := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/userinfo", config.AppConfig.Keycloak.URL, config.AppConfig.Keycloak.Realm)
	req, err := http.NewRequest("GET", keycloakUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var profile map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return nil, err
	}

	return profile, nil
}
