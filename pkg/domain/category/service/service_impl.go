package service

import (
	"github.com/Mangaba-Labs/ape-finance-api/pkg/api/models"
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/category/model"
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/category/repository"
)

// CategoryServiceImpl implementation of CategoryService
type CategoryServiceImpl struct {
	categoryRepository repository.CategoryRepository
}

// CreateCategory service implementation
func (c *CategoryServiceImpl) CreateCategory(category *model.Category) (apiResponse models.ApiResponse) {
	err := c.categoryRepository.Create(category)
	if err != nil {
		apiResponse.Set("Error", "Could not create category", 500)
	} else {
		apiResponse.Set("Success", "Created", 201)
	}
	return apiResponse
}

// DeleteCategory service implementation
func (c *CategoryServiceImpl) DeleteCategory(ID int) (apiResponse models.ApiResponse) {
	_, err := c.categoryRepository.FindByID(ID)
	if err != nil {
		apiResponse.Set("Error", "Category not found", 404)
	}
	err = c.categoryRepository.Delete(ID)
	if err != nil {
		apiResponse.Set("Error", "Could not delete category", 500)
	} else {
		apiResponse.Set("Success", "Deleted", 200)
	}
	return apiResponse
}

// EditCategory service implementation
func (c *CategoryServiceImpl) EditCategory(category *model.Category) error {
	return nil
}

// GetCategories service implementation
func (c *CategoryServiceImpl) GetCategories(ID int) (categoriesResponse []model.CategoryResponse, apiResponse models.ApiResponse) {
	categories, err := c.categoryRepository.FindAllByUser(ID)
	if err != nil {
		apiResponse.Set("Error", "Could not get your categories", 500)
	} else {
		categoriesResponse = parseAllCategories(categories)
		apiResponse.Set("Success", "Ok!", 200)
	}
	return categoriesResponse, apiResponse
}

func parseAllCategories(categories []model.Category) []model.CategoryResponse {
	categoriesResponse := []model.CategoryResponse{}
	for i := 0; i < len(categories); i++ {
		var category model.CategoryResponse
		category.ParseFromDatabase(categories[i])
		categoriesResponse = append(categoriesResponse, category)
	}
	return categoriesResponse
}
