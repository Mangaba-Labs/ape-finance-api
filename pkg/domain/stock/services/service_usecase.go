package services

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/stock/model"
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/stock/repository"
	"github.com/mxschmitt/playwright-go"
)

// Service struct implementation
type Service struct {
	Repository repository.Repository
}

// CreateStock search stock in tradingview and send to repository
func (s Service) CreateStock(stock *model.StockModel) error {
	err := s.Repository.Create(stock)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}

// GetStocks service to get all stocks in database
func (s Service) GetStocks(userID int) ([]model.StockResponse, error) {
	stockModels, err := s.Repository.FindAllByID(userID)
	if err != nil {
		return nil, err
	}
	// stockResponse := &model.StockResponse{}
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

// CheckIfExists check if stock exists before register in database
func (s Service) CheckIfExists(bvmf string) (stock model.StockModel, err error) {
	// It is too much faster check if the stock is already in database
	stock, err = s.Repository.FindByBvmf(bvmf)
	if err == nil {
		return
	}
	// Opening browser
	pw, err := playwright.Run()
	if err != nil {
		return
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		return
	}
	page, err := browser.NewPage()

	if err != nil {
		log.Fatalf("could not create page: %v\n", err)
		return
	}

	searchPage := fmt.Sprintf("https://www.tradingview.com/symbols/BMFBOVESPA-%s/", bvmf)

	if _, err = page.Goto(searchPage); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	// Finding company name
	companyEntry, err := page.QuerySelectorAll("div.tv-symbol-header__first-line")

	if err != nil {
		log.Fatalln(err)
		return
	}

	if len(companyEntry) == 0 {
		return stock, errors.New("Company not founded")
	}

	stock.Company, _ = companyEntry[0].InnerText()

	// Closing browser
	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v\n", err)
		return
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v\n", err)
		return
	}
	return stock, nil
}

func scrapStock(browser playwright.Browser, bvmf string) (scrapped model.VariableData, err error) {
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v\n", err)
		return
	}

	searchPage := fmt.Sprintf("https://www.tradingview.com/symbols/BMFBOVESPA-%s/", bvmf)
	if _, err = page.Goto(searchPage); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	// Variation
	variationValuesEntry, err := page.QuerySelectorAll("div.js-symbol-change-direction.tv-symbol-price-quote__change")
	if err != nil {
		log.Fatalln(err)
		return
	}
	variation, err := variationValuesEntry[0].InnerText()
	if err != nil {
		log.Fatalln(err)
		return
	}
	// Stock Value
	valueEntry, err := page.QuerySelectorAll("div.tv-symbol-price-quote__value.js-symbol-last")
	if err != nil {
		log.Fatalf("could not get entries: %v\n", err)
		return
	}
	value, err := valueEntry[0].InnerText()
	price, _ := strconv.ParseFloat(value, 2)

	scrapped.Price = float32(price)
	scrapped.Variation = variation
	return scrapped, nil
}

// Async method to get scrapped data and parse to stockResponse
func worker(wg *sync.WaitGroup, browser playwright.Browser, stock model.StockModel, stockResponse *[]model.StockResponse) {
	defer wg.Done()

	scrapped, _ := scrapStock(browser, stock.Bvmf)

	var response model.StockResponse

	response.ParseModelToResponse(stock, scrapped)

	*stockResponse = append(*stockResponse, response)
}
