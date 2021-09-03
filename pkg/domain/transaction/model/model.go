package model

import (
	"time"

	"gorm.io/gorm"
)

// Transaction to use in database
type Transaction struct {
	gorm.Model
	Date        time.Time `gorm:"not null" json:"date"`
	Description string    `gorm:"not null" json:"description"`
	IDCategory  uint64    `gorm:"not null; foreignKey:ID;references:categories" json:"id_category"`
	IDUser      uint64    `gorm:"not null; foreignKey:ID;references:User;"`
	Type        string    `gorm:"not null" json:"type"`
	Value       float32   `gorm:"not null" json:"value"`
}

// Transaction to use in database
type TransactionResponse struct {
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	ID          uint64    `json:"id"`
	IDCategory  uint64    `json:"id_category"`
	IDUser      uint64    `json:"id_user"`
	Type        string    `json:"type"`
	Value       float32   `json:"value"`
}

func (t *Transaction) IsValidFields() bool {
	if len(t.Description) < 1 || t.IDCategory < 1 || len(t.Type) < 1 || t.Value < 0 {
		return false
	}
	return true
}

func (t *TransactionResponse) ParseFromDatabase(transaction Transaction) {
	t.Date = transaction.Date
	t.Description = transaction.Description
	t.ID = uint64(transaction.ID)
	t.IDCategory = transaction.IDCategory
	t.IDUser = transaction.IDUser
	t.Type = transaction.Type
	t.Value = transaction.Value
}
