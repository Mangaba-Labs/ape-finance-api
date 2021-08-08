package handler

import (
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/category/model"
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/category/service"
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/utils"
	"github.com/gofiber/fiber/v2"
)

// CategoryHandlerImpl implementation of CategoryHandler
type CategoryHandlerImpl struct {
	service service.CategoryService
}

// CreateCategory implementation
func (h *CategoryHandlerImpl) CreateCategory(c *fiber.Ctx) error {
	var category = &model.Category{}
	if err := c.BodyParser(category); err != nil {
		return setupResponse(c, 400, "Error", "Bad Request!")
	}
	if len(category.Name) < 1 {
		return setupResponse(c, 400, "Error", "Category name cannot be null!")
	}
	category.IDUser = utils.GetUserID(c)

	response := h.service.CreateCategory(category)
	return setupResponse(c, response.HttpCode, response.Status, response.Message)
}

// GetCategories implementation
func (h *CategoryHandlerImpl) GetCategories(c *fiber.Ctx) error {
	categories, response := h.service.GetCategories(utils.GetUserID(c))
	return setupResponseWithItems(c, response.HttpCode, response.Status, response.Message, categories)
}

// EditCategory implementation
func (h *CategoryHandlerImpl) EditCategory(c *fiber.Ctx) error {
	return setupResponse(c, 200, "Success", "Ok!")
}

// DeleteCategory implementation
func (h *CategoryHandlerImpl) DeleteCategory(c *fiber.Ctx) error {
	response := h.service.DeleteCategory(utils.GetUserID(c))
	return setupResponse(c, response.HttpCode, response.Status, response.Message)
}

func setupResponse(c *fiber.Ctx, httpCode int, status string, message string) error {
	return c.Status(httpCode).JSON(fiber.Map{"status": status, "message": message})
}

func setupResponseWithItems(c *fiber.Ctx, httpCode int, status string, message string, categories []model.CategoryResponse) error {
	return c.Status(httpCode).JSON(fiber.Map{"status": status, "message": message, "categories": categories})
}
