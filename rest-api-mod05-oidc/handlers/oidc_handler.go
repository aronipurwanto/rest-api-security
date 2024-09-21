package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"rest-api-mod05-oidc/services"
)

// OAuthLogin mengarahkan pengguna ke penyedia identitas
func OAuthLogin(c *fiber.Ctx) error {
	url := services.OAuth2Config.AuthCodeURL("random-state-string")
	return c.Redirect(url)
}

// OAuthCallback menangani callback dari penyedia identitas setelah login berhasil
func OAuthCallback(c *fiber.Ctx) error {
	code := c.Query("code")
	if code == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Authorization code is missing")
	}

	// Tukarkan authorization code dengan token
	token, err := services.ExchangeCode(context.Background(), code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to exchange token: " + err.Error())
	}

	// Extract ID Token and Access Token
	idToken := token.Extra("id_token").(string)
	accessToken := token.AccessToken

	logrus.Infof("ID Token: %s", idToken)
	logrus.Infof("Access Token: %s", accessToken)

	return c.JSON(fiber.Map{
		"access_token": accessToken,
		"id_token":     idToken,
		"token_type":   token.TokenType,
		"expiry":       token.Expiry,
	})
}
