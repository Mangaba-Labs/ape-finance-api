package config

import (
	"log"

	expenseModel "github.com/Mangaba-Labs/ape-finance-api/pkg/domain/expense/model"
	incomeModel "github.com/Mangaba-Labs/ape-finance-api/pkg/domain/income/model"
	stockModel "github.com/Mangaba-Labs/ape-finance-api/pkg/domain/stock/model"
	userModel "github.com/Mangaba-Labs/ape-finance-api/pkg/domain/user/model"
	"gorm.io/gorm"
)

// Migrate struct for database
type Migrate struct {
	DB *gorm.DB
}

// MigrateAll database migration
func (m *Migrate) MigrateAll() (err error) {
	log.Println("Migrating database... 🤞")
	err = m.DB.AutoMigrate(&userModel.User{}, &stockModel.StockModel{}, &expenseModel.ExpenseModel{}, &expenseModel.ExpenseCategories{}, &incomeModel.IncomeModel{}, &incomeModel.IncomeCategories{})

	if err != nil {
		log.Fatal("Something went wrong on db migration process...\n ", err)
	}

	log.Println("Database migrated with success 😁")
	return
}
