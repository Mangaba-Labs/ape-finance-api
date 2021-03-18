package handler

import (
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/stock/model"
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/stock/services"
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/utils"
	"github.com/gofiber/fiber/v2"
)

// Handler for stock service
type Handler struct {
	service services.StockService
}

// CreateStock handler for POST /stock
func (h *Handler) CreateStock(c *fiber.Ctx) error {
	var stock = &model.StockModel{}
	if err := c.BodyParser(stock); err != nil {
		return c.Status(422).JSON(fiber.Map{"status": "error", "message": err})
	}

	if stock.BuyPrice <= 0 {
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": "buy_price cannot be null"})
	}
	if stock.Quantity <= 0 {
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": "quantity cannot be less than one"})
	}

	identity := utils.GetTokenValue(c, "identity")
	stockInDB, err := h.service.CheckIfExists(stock.Bvmf)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Stock doesn't exist!"})
	}

	/*
		Here our interface{} need to be casted to float and after that to int
		if we do a .(int) directly will result in error:
		interface conversion: interface {} is float64, not int
	*/
	stock.IDUser = int(identity.(float64))
	stock.Company = stockInDB.Company
	err = h.service.CreateStock(stock)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err})
	}
	return nil
}
