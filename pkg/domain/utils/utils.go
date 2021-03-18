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
