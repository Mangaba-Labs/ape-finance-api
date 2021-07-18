package service

import (
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/category/model"
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/category/repository"
)

// CategoryServiceImpl implementation of CategoryService
type CategoryServiceImpl struct {
	categoryRepository repository.CategoryRepository
}

// CreateCategory service implementation
func (c *CategoryServiceImpl) CreateCategory(category *model.Category) error {
	return nil
}

// DeleteCategory service implementation
func (c *CategoryServiceImpl) DeleteCategory(ID int) error {
	return nil
}

// EditCategory service implementation
func (c *CategoryServiceImpl) EditCategory(category *model.Category) error {
	return nil
}

// GetCategories service implementation
func (c *CategoryServiceImpl) GetCategories() ([]model.Category, error) {
	return nil, nil
}
