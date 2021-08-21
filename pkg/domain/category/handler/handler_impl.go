package handler

import (
	"errors"
	"strconv"

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
	field, err := checkCategoryFields(category)
	if err != nil {
		return setupResponse(c, 400, "Error", "Category "+field+"cannot be null!")
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
	var category = &model.Category{}

	if err := c.BodyParser(category); err != nil {
		return setupResponse(c, 400, "Error", "Bad Request!")
	}
	field, err := checkCategoryFields(category)
	if category.ID < 1 || category.IDUser < 1 {
		field = "id"
	}
	if err != nil {
		return setupResponse(c, 400, "Error", "Category "+field+"cannot be null!")
	}

	response := h.service.EditCategory(category)
	return setupResponse(c, response.HttpCode, response.Status, response.Message)
}

// DeleteCategory implementation
func (h *CategoryHandlerImpl) DeleteCategory(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return setupResponse(c, 400, "error", "Invalid parameters!")
	}
	response := h.service.DeleteCategory(uint(id))
	return setupResponse(c, response.HttpCode, response.Status, response.Message)
}

func setupResponse(c *fiber.Ctx, httpCode int, status string, message string) error {
	return c.Status(httpCode).JSON(fiber.Map{"status": status, "message": message})
}

func setupResponseWithItems(c *fiber.Ctx, httpCode int, status string, message string, categories []model.CategoryResponse) error {
	return c.Status(httpCode).JSON(fiber.Map{"status": status, "message": message, "categories": categories})
}

func checkCategoryFields(category *model.Category) (string, error) {
	if len(category.Name) < 1 {
		return "category", errors.New("invalid fields")
	}
	if len(category.Type) < 1 {
		return "type", errors.New("invalid fields")
	}
	return "", nil
}
