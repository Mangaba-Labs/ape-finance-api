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
func (c *CategoryRepositoryImpl) Delete(ID uint) error {
	result := c.DB.Where("id = ?", ID).Delete(&model.Category{})
	return result.Error
}

// Edit repository implementation in category PUT
func (c *CategoryRepositoryImpl) Edit(category *model.Category) error {
	var oldCategory model.Category
	c.DB.Where("id = ?", category.ID).Find(&oldCategory)
	oldCategory.Name = category.Name
	oldCategory.Type = category.Type
	result := c.DB.Save(&oldCategory)
	return result.Error
}

// FindByID repository implementation
func (c *CategoryRepositoryImpl) FindByID(ID uint) (category model.Category, err error) {
	result := c.DB.Find(&category).Where("id = ?", ID)
	return category, result.Error
}

// FindAllByUser repository implementation in GET categories
func (c *CategoryRepositoryImpl) FindAllByUser(ID uint64) (categories []model.Category, err error) {
	result := c.DB.Find(&categories).Where("id_user = ?", ID)
	return categories, result.Error
}
