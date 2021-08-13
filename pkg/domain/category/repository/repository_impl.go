package repository

import (
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/category/model"
	"gorm.io/gorm"
)

// CategoryRepositoryImpl struct implementation of CategoryRepository
type CategoryRepositoryImpl struct {
	DB *gorm.DB
}

// Create repository implementation in category POST/creation
func (c *CategoryRepositoryImpl) Create(category *model.Category) error {
	result := c.DB.Create(&category)
	return result.Error
}

// Delete repository implementation in category DELETE
func (c *CategoryRepositoryImpl) Delete(ID string) error {
	result := c.DB.Where("id = ?", ID).Delete(&model.Category{})
	return result.Error
}

// Edit repository implementation in category PUT
func (c *CategoryRepositoryImpl) Edit(*model.Category) error {
	return nil
}

// FindByID repository implementation
func (c *CategoryRepositoryImpl) FindByID(ID string) (category model.Category, err error) {
	result := c.DB.Find(&category).Where("id = ?", ID)
	return category, result.Error
}

// FindAllByUser repository implementation in GET categories
func (c *CategoryRepositoryImpl) FindAllByUser(ID int) (categories []model.Category, err error) {
	result := c.DB.Find(&categories).Where("id_user = ?", ID)
	return categories, result.Error
}
