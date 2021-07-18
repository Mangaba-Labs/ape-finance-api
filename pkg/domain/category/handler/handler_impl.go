package handler

import (
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/category/service"
	"github.com/gofiber/fiber/v2"
)

// CategoryHandlerImpl implementation of CategoryHandler
type CategoryHandlerImpl struct {
	service service.CategoryService
}

// CreateCategory implementation
func (h *CategoryHandlerImpl) CreateCategory(c *fiber.Ctx) error {
	return nil
}

// GetCategories implementation
func (h *CategoryHandlerImpl) GetCategories(c *fiber.Ctx) error {
	return nil
}

// EditCategory implementation
func (h *CategoryHandlerImpl) EditCategory(c *fiber.Ctx) error {
	return nil
}

// DeleteCategory implementation
func (h *CategoryHandlerImpl) DeleteCategory(c *fiber.Ctx) error {
	return nil
}
