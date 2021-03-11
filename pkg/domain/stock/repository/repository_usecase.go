package repository

import (
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/stock/model"
	"gorm.io/gorm"
)

// Repository concrete type
type Repository struct {
	DB *gorm.DB
}

// FindAllByID find all user's stocks by ID
func (r Repository) FindAllByID(id int) ([]model.StockModel, error) {
	return nil, nil
}

// Create stock in database
func (r Repository) Create(stock model.StockModel) error {
	return nil
}

// Delete stock in database
func (r Repository) Delete(id int) error {
	return nil
}
