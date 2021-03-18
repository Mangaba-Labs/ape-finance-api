package repository

import (
	"errors"

	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/stock/model"
	"gorm.io/gorm"
)

// Repository concrete type
type Repository struct {
	DB *gorm.DB
}

// FindAllByID find all user's stocks by ID
func (r Repository) FindAllByID(id int) (stocks []model.StockModel, err error) {
	return
}

// Create stock in database
func (r Repository) Create(stock *model.StockModel) (err error) {
	result := r.DB.Create(&stock)
	err = result.Error
	rowsCount := result.RowsAffected
	if err != nil || rowsCount <= 0 {
		return errors.New("Cannot create stock")
	}
	return
}

// Delete stock in database
func (r Repository) Delete(id int) (err error) {
	return
}

// FindByBvmf to check if stock is already in our database before use our crawler
func (r Repository) FindByBvmf(bvmf string) (stock model.StockModel, err error) {
	result := r.DB.First(&stock, "bvmf = ?", bvmf)
	err = result.Error
	return
}
