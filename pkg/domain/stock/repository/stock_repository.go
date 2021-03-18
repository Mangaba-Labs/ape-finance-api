package repository

import (
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/stock/model"
	"gorm.io/gorm"
)

// StockRepository Contract
type StockRepository interface {
	FindAllByID(id int) ([]model.StockModel, error)
	Create(*model.StockModel) error
	Delete(id int) error
}

// NewStockRepository repository postgres implementation
func NewStockRepository(db *gorm.DB) StockRepository {
	return &Repository{
		DB: db,
	}
}
