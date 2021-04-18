package model

import (
	"time"

	"gorm.io/gorm"
)

// IncomeModel to use in database
type IncomeModel struct {
	gorm.Model
	Description string    `gorm:"not null" json:"description"`
	Date        time.Time `gorm:"not null" json:"date"`
	IDUser      uint64    `gorm:"not null; foreignKey:ID;references:User;"`
	IDCategory  uint64    `gorm:"not null;foreignKey:ID;references:IncomeCategories"`
	Value       float32   `gorm:"not null" json:"value"`
}

// IncomeCategories default
type IncomeCategories struct {
	gorm.Model
	Name string `gorm:"not null"`
}
