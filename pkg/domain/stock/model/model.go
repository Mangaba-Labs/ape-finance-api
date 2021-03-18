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
	Variation string  `json:"variation"`
}

// ParseModelToResponse from scrapStock
func (s *StockResponse) ParseModelToResponse(stock StockModel, priceNow float32, variation string) {
	s.BuyPrice = stock.BuyPrice
	s.Bvmf = stock.Bvmf
	// s.ID = int(stock.ID)
	s.NowPrice = priceNow
	s.Variation = variation
}
