package config

import (
	"log"

	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/user"
	"gorm.io/gorm"
)

type Migrate struct {
	DB *gorm.DB
}

func (m *Migrate) MigrateAll() (err error) {
	log.Println("Migrating database... 🤞")
	err = m.DB.AutoMigrate(&user.User{})

	if err != nil {
		log.Fatal("Something went wrong on db migration process...\n ", err)
	}

	log.Println("Database migrated with success 😁")
	return
}
