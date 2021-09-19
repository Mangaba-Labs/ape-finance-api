package services

import (
	"log"
	"sync"

	"github.com/Mangaba-Labs/ape-finance-api/pkg/api/models"
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/stock/model"
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/stock/repository"
	"github.com/mxschmitt/playwright-go"
)

// Service struct implementation
type StockServiceImpl struct {
	Repository repository.Repository
}

// CreateStock search stock in tradingview and send to repository
func (s *StockServiceImpl) CreateStock(stock *model.StockModel) (apiResponse models.ApiResponse) {
	err := s.Repository.Create(stock)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}

// GetStocks service to get all stocks in database
func (s *StockServiceImpl) GetStocks(userID int) (stocks []model.StockResponse, apiResponse models.ApiResponse) {
	stockModels, err := s.Repository.FindAllByID(userID)
	if err != nil {
		return nil, err
	}
	pw, err := playwright.Run()
	if err != nil {
		return nil, err
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		return nil, err
	}

	responseSlice := []model.StockResponse{}

	var wg sync.WaitGroup
	for i := 0; i < len(stockModels); i++ {
		wg.Add(1)
		go worker(&wg, browser, stockModels[i], &responseSlice)
	}
	wg.Wait()
	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v\n", err)
		return nil, err
	}
	return responseSlice, nil
}

func isInDatabase(bvmf string) bool {
	return false
}
