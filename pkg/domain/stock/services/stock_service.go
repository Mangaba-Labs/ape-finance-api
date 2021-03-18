package services

import (
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/stock/model"
	stockRepository "github.com/Mangaba-Labs/ape-finance-api/pkg/domain/stock/repository"
)

// StockService contract
type StockService interface {
	CreateStock(stock *model.StockModel) error
	CheckIfExists(bvmf string) (model.StockModel, error)
}

// NewUserService returns a UserService implementation
func NewUserService(repository stockRepository.Repository) (service StockService) {
	service = &Service{
		Repository: repository,
	}
	return
}
