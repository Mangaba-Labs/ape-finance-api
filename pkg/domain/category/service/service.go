package service

import (
	"github.com/Mangaba-Labs/ape-finance-api/pkg/api/models"
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/category/model"
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/category/repository"
)

// CategoryService contract
type CategoryService interface {
	CreateCategory(category *model.Category) models.ApiResponse
	DeleteCategory(ID int) models.ApiResponse
	EditCategory(category *model.Category) error
	GetCategories(ID int) ([]model.CategoryResponse, models.ApiResponse)
}

// NewCategoryService constructor
func NewCategoryService(repository repository.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{
		categoryRepository: repository,
	}
}
