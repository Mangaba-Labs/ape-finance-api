package handler

import (
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/stock/model"
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/stock/services"
	"github.com/gofiber/fiber/v2"
)

// StockHandler interface
type StockHandler interface {
	Post(c *fiber.Ctx) error
	Get(c *fiber.Ctx) ([]model.StockResponse, error)
	Put(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

// NewStockHandler returns a pointer to an handler impl
func NewStockHandler(s services.StockService) Handler {
	return StockHandlerImpl{
		service: s,
	}
}
