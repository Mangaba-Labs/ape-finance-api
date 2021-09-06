package repository

import (
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/transaction/model"
	"gorm.io/gorm"
)

// TransactionRepository contract
type TransactionRepository interface {
	Create(*model.Transaction) error
	Delete(ID uint) error
	FindByID(ID uint) (model.Transaction, error)
	FindAllByUser(ID uint64) ([]model.Transaction, error)
	Update(*model.Transaction) error
}

// NewTransactionRepository constructor
func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &TransactionRepositoryImpl{
		DB: db,
	}
}
