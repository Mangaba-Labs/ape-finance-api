package model

import (
	"time"

	"gorm.io/gorm"
)

// ExpenseModel to use in database
type ExpenseModel struct {
	gorm.Model
	Description string    `gorm:"not null" json:"description"`
	Date        time.Time `gorm:"not null" json:"date"`
	IDUser      uint64    `gorm:"not null; foreignKey:ID;references:User;"`
	IDCategory  uint64    `gorm:"not null;foreignKey:ID;references:ExpenseCategories"`
	Value       float32   `gorm:"not null" json:"value"`
}

// ExpenseCategories default
type ExpenseCategories struct {
	gorm.Model
	Name string `gorm:"not null"`
}
