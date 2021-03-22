package model

import "gorm.io/gorm"

// StockModel to save in database
type StockModel struct {
	gorm.Model
	BuyPrice float32 `gorm:"not null" json:"buy_price"`
	Bvmf     string  `gorm:"not null" json:"bvmf"`
	Company  string  `gorm:"not null" json:"company"`
	IDUser   int     `gorm:"not null; foreignKey:ID;references:User;"`
	Quantity int     `gorm:"not null" json:"quantity"`
}

// StockResponse to endpoints response
type StockResponse struct {
	BuyPrice  float32 `json:"buy_price"`
	Bvmf      string  `json:"bvmf"`
	Company   string  `json:"company"`
	NowPrice  float32 `json:"now_price"`
	Quantity  int     `json:"quantity"`
	Variation string  `json:"variation"`
	Image     string  `json:"image"`
}

type VariableData struct {
	Variation string
	Price     float32
	Image     string
}

// ParseModelToResponse from scrapStock
func (s *StockResponse) ParseModelToResponse(stock StockModel, scrappedData VariableData) {
	s.BuyPrice = stock.BuyPrice
	s.Bvmf = stock.Bvmf
	s.NowPrice = scrappedData.Price
	s.Variation = scrappedData.Variation
	s.Quantity = stock.Quantity
	s.Company = stock.Company
	s.Image = scrappedData.Image
}
