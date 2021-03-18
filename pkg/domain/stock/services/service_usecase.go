package services

import (
	"errors"
	"fmt"
	"log"
	"strconv"

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

// CheckIfExists check if stock exists before register in database
func (s Service) CheckIfExists(bvmf string) (stock model.StockModel, err error) {
	// It is too much faster check if the stock is already in database
	stock, err = s.Repository.FindByBvmf(bvmf)
	if err == nil {
		return
	}
	// Opening browser
	pw, err := playwright.Run()
	// Could not start playwright
	if err != nil {
		return
	}
	browser, err := pw.Chromium.Launch()
	// Could not launch browser
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

func scrapStock(browser playwright.Browser, stock model.StockModel) (model.StockModel, error) {
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v\n", err)
		return stock, err
	}

	searchPage := fmt.Sprintf("https://www.tradingview.com/symbols/BMFBOVESPA-%s/", stock.Bvmf)

	if _, err = page.Goto(searchPage); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	variationValuesEntry, err := page.QuerySelectorAll("div.js-symbol-change-direction.tv-symbol-price-quote__change")

	if err != nil {
		log.Fatalln(err)
		return stock, err
	}

	variationValue, err := variationValuesEntry[0].InnerText()

	if err != nil {
		log.Fatalln(err)
		return stock, err
	}

	nameEntry, err := page.QuerySelectorAll("div.tv-symbol-header__first-line")

	if err != nil {
		log.Fatalln(err)
		return stock, err
	}

	name, err := nameEntry[0].InnerText()

	if err != nil {
		log.Fatalln(err)
		return stock, err
	}

	valueEntry, err := page.QuerySelectorAll("div.tv-symbol-price-quote__value.js-symbol-last")
	if err != nil {
		log.Fatalf("could not get entries: %v\n", err)
		return stock, err
	}

	// Share Value
	value, err := valueEntry[0].InnerText()

	valueFloat, _ := strconv.ParseFloat(value, 2)

	fmt.Printf("Company: %s, Price: R$%.2f, Variation: %s\n", name, valueFloat, variationValue)

	return stock, nil
}
