package handler

import (
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/category/service"
	"github.com/gofiber/fiber/v2"
)

// CategoryHandler contract
type CategoryHandler interface {
	CreateCategory(c *fiber.Ctx) error
	GetCategories(c *fiber.Ctx) error
	EditCategory(c *fiber.Ctx) error
	DeleteCategory(c *fiber.Ctx) error
}

// NewCategoryHandler constructor
func NewCategoryHandler(service service.CategoryService) CategoryHandler {
	return &CategoryHandlerImpl{
		service: service,
	}
}
