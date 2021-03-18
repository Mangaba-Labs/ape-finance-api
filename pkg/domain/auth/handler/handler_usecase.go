package handler

import (
	"os"
	"time"

	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/user/model"
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/user/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandlerUsecase struct {
	s services.UserService
}

func (a *AuthHandlerUsecase) Login(c *fiber.Ctx) error {
	var input model.AuthRequest
	if err := c.BodyParser(&input); err != nil {
		return c.JSON(fiber.Map{"status": "error", "error": "malformed auth request", "data": nil})

	}
	email := input.Email
	pass := input.Password

	usr, err := a.s.GetUserByEmail(email)

	if err != nil || len(usr.Email) == 0 {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(pass)); err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["identity"] = usr.ID
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 168).Unix()

	t, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
}
