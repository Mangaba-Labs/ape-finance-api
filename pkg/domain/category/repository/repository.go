package repository

import (
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/category/model"
	"gorm.io/gorm"
)

// CategoryRepository contract
type CategoryRepository interface {
	Create(*model.Category) error
	Delete(ID int) error
	Edit(*model.Category) error
	FindAllByUser(ID int) ([]model.Category, error)
}

// NewCategoryRepository constructor
func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{
		DB: db,
	}
}
