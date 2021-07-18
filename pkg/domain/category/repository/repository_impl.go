package repository

import (
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/category/model"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	DB *gorm.DB
}

// Create repository implementation in category POST/creation
func (c *CategoryRepositoryImpl) Create(*model.Category) error {
	return nil
}

// Delete repository implementation in category DELETE
func (c *CategoryRepositoryImpl) Delete(ID int) error {
	return nil
}

// Edit repository implementation in category PUT
func (c *CategoryRepositoryImpl) Edit(*model.Category) error {
	return nil
}

// FindAllByUser repository implementation in GET categories
func (c *CategoryRepositoryImpl) FindAllByUser(ID int) (categories []model.Category, err error) {
	return nil, nil
}
