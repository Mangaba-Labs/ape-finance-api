package service

import (
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/category/model"
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/category/repository"
)

// CategoryService contract
type CategoryService interface {
	CreateCategory(category *model.Category) error
	DeleteCategory(ID int) error
	EditCategory(category *model.Category) error
	GetCategories() ([]model.Category, error)
}

// NewCategoryService constructor
func NewCategoryService(repository repository.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{
		categoryRepository: repository,
	}
}
