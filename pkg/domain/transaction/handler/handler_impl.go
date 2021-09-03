package handler

import (
	"strconv"

	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/transaction/model"
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/transaction/service"
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/utils"
	"github.com/gofiber/fiber/v2"
)

type TransactionHandlerImpl struct {
	service service.TransactionService
}

func (i *TransactionHandlerImpl) Add(c *fiber.Ctx) error {
	var transaction = &model.Transaction{}
	if err := c.BodyParser(transaction); err != nil {
		return setupResponse(c, 400, "error", "Bad Request!")
	}
	if !transaction.IsValidFields() {
		return setupResponse(c, 400, "error", "Bad Request!")
	}
	transaction.IDUser = utils.GetUserID(c)
	response := i.service.Create(transaction)
	return setupResponse(c, response.HttpCode, response.Status, response.Message)
}

func (i *TransactionHandlerImpl) FindAll(c *fiber.Ctx) error {
	itens, response := i.service.GetAllByUser(utils.GetUserID(c))
	return setupResponseWithItems(c, response.HttpCode, response.Status, response.Message, itens)
}

func (i *TransactionHandlerImpl) Edit(c *fiber.Ctx) error {
	var transaction = &model.Transaction{}
	if err := c.BodyParser(transaction); err != nil {
		return setupResponse(c, 400, "error", "Bad Request!")
	}
	return setupResponse(c, 200, "success", "Nothing happens!")
}

func (i *TransactionHandlerImpl) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return setupResponse(c, 400, "error", "Invalid Parameters!")
	}
	response := i.service.Delete(uint(id))
	return setupResponse(c, response.HttpCode, response.Status, response.Message)
}

func setupResponse(c *fiber.Ctx, httpCode int, status string, message string) error {
	return c.Status(httpCode).JSON(fiber.Map{"status": status, "message": message})
}

func setupResponseWithItems(c *fiber.Ctx, httpCode int, status string, message string, categories []model.TransactionResponse) error {
	return c.Status(httpCode).JSON(fiber.Map{"status": status, "message": message, "categories": categories})
}
