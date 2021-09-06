package service

import (
	"github.com/Mangaba-Labs/ape-finance-api/pkg/api/models"
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/transaction/model"
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/transaction/repository"
)

// TransactionService contract
type TransactionService interface {
	Create(transaction *model.Transaction) models.ApiResponse
	Delete(ID uint) models.ApiResponse
	Edit(transaction *model.Transaction) models.ApiResponse
	GetAllByUser(ID uint64) ([]model.TransactionResponse, models.ApiResponse)
}

// NewTransactionService constructor
func NewTransactionService(repository repository.TransactionRepository) TransactionService {
	return &TransactionServiceImpl{
		transactionRepository: repository,
	}
}
