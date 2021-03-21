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
}

type VariableData struct {
	Variation string
	Price     float32
}

// ParseModelToResponse from scrapStock
func (s *StockResponse) ParseModelToResponse(stock StockModel, scrappedData VariableData) {
	s.BuyPrice = stock.BuyPrice
	s.Bvmf = stock.Bvmf
	s.NowPrice = scrappedData.Price
	s.Variation = scrappedData.Variation
	s.Quantity = stock.Quantity
}

// -> Abrir o browser
// --> Loop chamando os workers, cada workers será uma stock
// --> Cada worker abrirá uma página no mesmo browser

// -> Fechar browser Após terminar os workers

// -> Parse de model para response
