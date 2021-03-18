package handler

import (
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/stock/services"
	"github.com/gofiber/fiber/v2"
)

// StockHandler interface
type StockHandler interface {
	CreateStock(c *fiber.Ctx) error
}

// NewStockHandler returns a pointer to an handler impl
func NewStockHandler(s services.StockService) Handler {
	return Handler{
		service: s,
	}
}
