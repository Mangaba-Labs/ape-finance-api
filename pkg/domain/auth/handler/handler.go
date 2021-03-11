package handler

import (
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/user/services"
	"github.com/gofiber/fiber/v2"
)

// AuthHandler describe the auth resource
type AuthHandler interface {
	Login(ctx *fiber.Ctx) error
}

// NewAuthHandler returns a pointer to an handler impl
func NewAuthHandler(s services.UserService) AuthHandler {
	return &AuthHandlerUsecase{
		s: s,
	}
}
