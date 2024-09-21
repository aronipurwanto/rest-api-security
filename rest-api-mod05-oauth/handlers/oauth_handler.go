package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"rest-api-mod05-oauth/services"
)

// Redirect user ke OAuth provider
func OAuthLogin(c *fiber.Ctx) error {
	url := services.OAuth2Config.AuthCodeURL("random-state-string")
	return c.Redirect(url)
}

// Callback dari OAuth provider
func OAuthCallback(c *fiber.Ctx) error {
	code := c.Query("code")
	if code == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Authorization code is missing")
	}

	// Exchange authorization code for access token
	token, err := services.ExchangeCode(context.Background(), code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to exchange token: " + err.Error())
	}

	logrus.Infof("Received token: %+v", token)
	return c.JSON(fiber.Map{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
		"token_type":    token.TokenType,
		"expiry":        token.Expiry,
	})
}
