package handler

import (
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/transaction/service"
	"github.com/gofiber/fiber/v2"
)

// TransactionHandler contract
type TransactionHandler interface {
	Add(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
	Edit(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

// NewTransactionHandler constructor
func NewTransactionHandler(service service.TransactionService) TransactionHandler {
	return &TransactionHandlerImpl{
		service: service,
	}
}
