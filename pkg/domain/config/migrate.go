package config

import (
	"log"

	category "github.com/Mangaba-Labs/ape-finance-api/pkg/domain/category/model"
	stock "github.com/Mangaba-Labs/ape-finance-api/pkg/domain/stock/model"
	transaction "github.com/Mangaba-Labs/ape-finance-api/pkg/domain/transaction/model"
	user "github.com/Mangaba-Labs/ape-finance-api/pkg/domain/user/model"
	"gorm.io/gorm"
)

// Migrate struct for database
type Migrate struct {
	DB *gorm.DB
}

// MigrateAll database migration
func (m *Migrate) MigrateAll() (err error) {
	log.Println("Migrating database... 🤞")
	err = m.DB.AutoMigrate(&user.User{}, &stock.StockModel{}, &transaction.Transaction{}, &category.Category{})

	if err != nil {
		log.Fatal("Something went wrong on db migration process...\n ", err)
	}

	log.Println("Database migrated with success 😁")
	return
}
