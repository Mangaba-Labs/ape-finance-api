package repository

import (
	"errors"

	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/stock/model"
	"gorm.io/gorm"
)

// Repository concrete type
type StockRepositoryImpl struct {
	DB *gorm.DB
}

// FindAllByID find all user's stocks by ID
func (r StockRepositoryImpl) FindAllByID(id int) (stocks []model.StockModel, err error) {
	result := r.DB.Where("id_user = ?", id).Find(&stocks)
	err = result.Error
	if err != nil {
		return nil, err
	}
	return stocks, nil
}

// Create stock in database
func (r StockRepositoryImpl) Create(stock *model.StockModel) (err error) {
	result := r.DB.Create(&stock)
	err = result.Error
	rowsCount := result.RowsAffected
	if err != nil || rowsCount <= 0 {
		return errors.New("Cannot create stock")
	}
	return
}

// Delete stock in database
func (r StockRepositoryImpl) Delete(id int) (err error) {
	result := r.DB.Where("id = ?", id).Delete(&model.StockModel{})
	return result.Error
}

// FindByBvmf to check if stock is already in our database before use our crawler
func (r StockRepositoryImpl) findByBvmf(bvmf string) (stock model.StockModel, err error) {
	result := r.DB.First(&stock, "bvmf = ?", bvmf)
	err = result.Error
	return
}
