package controller

import (
	"github.com/gofiber/fiber/v2"
	"rest-api-pos/model"
	"rest-api-pos/response"
)

func GetUser(c *fiber.Ctx) error {
	// Contoh data pengguna
	user := model.User{
		ID:    1,
		Name:  "Roni Purwanto",
		Email: "roni@example.com",
		Roles: []string{"admin", "user"},
	}
	return c.JSON(response.SuccessResponse(user, "User fetched successfully"))
}
