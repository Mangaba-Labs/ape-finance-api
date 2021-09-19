package services

import (
	"github.com/Mangaba-Labs/ape-finance-api/pkg/api/models"
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/stock/model"
	stockRepository "github.com/Mangaba-Labs/ape-finance-api/pkg/domain/stock/repository"
)

// StockService contract
type StockService interface {
	CreateStock(stock *model.StockModel) models.ApiResponse
	GetStocks(userID int) ([]model.StockResponse, models.ApiResponse)
}

// NewStockService returns a NewStockService implementation
func NewStockService(repository stockRepository.Repository) (service StockService) {
	return &StockServiceImpl{
		Repository: repository,
	}
}
