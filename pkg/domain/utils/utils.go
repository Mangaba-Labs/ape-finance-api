package utils

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// GetTokenValue in Authorization token
func GetTokenValue(c *fiber.Ctx, key string) interface{} {
	decoded := c.Locals("user").(*jwt.Token)
	userInfo := decoded.Claims.(jwt.MapClaims)
	value := userInfo[key]
	return value
}

// GetUserID from Authroization token
func GetUserID(c *fiber.Ctx) int {
	decoded := c.Locals("user").(*jwt.Token)
	userInfo := decoded.Claims.(jwt.MapClaims)
	value := userInfo["identity"]
	return int(value.(float64))
}
