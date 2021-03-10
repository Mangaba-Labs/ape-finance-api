package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwt "github.com/gofiber/jwt/v2"
	"os"
)

func Protected() fiber.Handler {
	secret := os.Getenv("TOKEN_SECRET")

	return jwt.New(jwt.Config{
		SigningKey:   []byte(secret),
		ErrorHandler: jwtErrorHandler,
	})
}

func jwtErrorHandler(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{"status": "error", "message": err.Error(), "data": nil})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}
