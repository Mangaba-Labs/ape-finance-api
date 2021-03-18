package config

import (
	"log"

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
	// err = m.DB.AutoMigrate(&user.User{})
	err = m.DB.AutoMigrate(&userModel.User{}, &stockModel.StockModel{})

	if err != nil {
		log.Fatal("Something went wrong on db migration process...\n ", err)
	}

	log.Println("Database migrated with success 😁")
	return
}
