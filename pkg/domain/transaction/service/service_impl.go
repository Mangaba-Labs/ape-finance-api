package service

import (
	"github.com/Mangaba-Labs/ape-finance-api/pkg/api/models"
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/transaction/model"
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/transaction/repository"
)

type TransactionServiceImpl struct {
	transactionRepository repository.TransactionRepository
}

func (t *TransactionServiceImpl) Create(transaction *model.Transaction) (apiResponse models.ApiResponse) {
	err := t.transactionRepository.Create(transaction)
	if err != nil {
		apiResponse.Set("Error", "Could not create transaction", 500)
	} else {
		apiResponse.Set("Success", "Created", 201)
	}
	return apiResponse
}

func (t *TransactionServiceImpl) Delete(ID uint) (apiResponse models.ApiResponse) {
	_, err := t.transactionRepository.FindByID(ID)
	if err != nil {
		apiResponse.Set("Error", "Transaction not found", 404)
	}
	err = t.transactionRepository.Delete(ID)
	if err != nil {
		apiResponse.Set("Error", "Could not delete transaction", 500)
	} else {
		apiResponse.Set("Success", "Deleted", 200)
	}
	return apiResponse
}

func (t *TransactionServiceImpl) Edit(transaction *model.Transaction) models.ApiResponse {
	return models.ApiResponse{}
}

func (t *TransactionServiceImpl) GetAllByUser(ID uint64) (transactionsResponse []model.TransactionResponse, apiResponse models.ApiResponse) {
	transactions, err := t.transactionRepository.FindAllByUser(ID)
	if err != nil {
		apiResponse.Set("Error", "Could not get your transactions", 500)
	} else {
		transactionsResponse = parseAllTransactions(transactions)
		apiResponse.Set("Success", "OK!", 200)
	}
	return transactionsResponse, apiResponse
}

func parseAllTransactions(transactions []model.Transaction) []model.TransactionResponse {
	transactionResponse := []model.TransactionResponse{}
	for i := 0; i < len(transactions); i++ {
		var transaction model.TransactionResponse
		transaction.ParseFromDatabase(transactions[i])
		transactionResponse = append(transactionResponse, transaction)
	}
	return transactionResponse
}
