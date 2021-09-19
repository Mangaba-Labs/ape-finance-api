package handler

import (
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/stock/model"
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/stock/services"
	"github.com/gofiber/fiber/v2"
)

// StockHandlerImpl for stock service
type StockHandlerImpl struct {
	service services.StockService
}

// CreateStock handler for POST /stock
func (s *StockHandlerImpl) CreateStock(c *fiber.Ctx) error {
	var stock = &model.StockModel{}
	if err := c.BodyParser(stock); err != nil {
		return c.Status(422).JSON(fiber.Map{"status": "error", "message": err})
	}
	// response := s.service.CreateStock(stock)
	// return setupResponse(c, response.)
}

// GetStocks get all regiters in database
func (s *StockHandlerImpl) GetStocks(c *fiber.Ctx) error {
	// identity := utils.GetTokenValue(c, "identity")
	// userID := int(identity.(float64))
	// response := s.service.GetStocks(userID)
	return nil
}

func setupResponse(c *fiber.Ctx, httpCode int, status string, message string) error {
	return c.Status(httpCode).JSON(fiber.Map{"status": status, "message": message})
}

func setupResponseWithItems(c *fiber.Ctx, httpCode int, status string, message string, stocks []model.StockResponse) error {
	return c.Status(httpCode).JSON(fiber.Map{"status": status, "message": message, "stocks": stocks})
}
