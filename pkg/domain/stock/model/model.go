package model

import "gorm.io/gorm"

// StockModel to save in database
type StockModel struct {
	gorm.Model
	BuyPrice float32 `json:"buy_price"`
	Bvmf     string  `json:"bvmf"` // ex: petr4
	Company  string  `json:"company"`
}

// StockResponse to endpoints response
type StockResponse struct {
	ID        int
	BuyPrice  float32 `json:"buy_price"`
	Bvmf      string  `json:"bvmf"`
	Company   string  `json:"company"`
	NowPrice  float32 `json:"now_price"`
	Variation string  `json:"variation"`
}
