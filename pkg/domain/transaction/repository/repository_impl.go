package repository

import (
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/transaction/model"
	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
	DB *gorm.DB
}

// Create transaction in database
func (t *TransactionRepositoryImpl) Create(transaction *model.Transaction) error {
	result := t.DB.Create(&transaction)
	return result.Error
}

// Delete transaction in database
func (t *TransactionRepositoryImpl) Delete(ID uint) error {
	result := t.DB.Where("id = ?", ID).Delete(&model.Transaction{})
	return result.Error
}

// FindByID transaction in database
func (t *TransactionRepositoryImpl) FindByID(ID uint) (model.Transaction, error) {
	var transaction model.Transaction
	result := t.DB.Find(&transaction).Where("id = ?", ID)
	return transaction, result.Error
}

// FindAllByUser finds all transactions in database
func (t *TransactionRepositoryImpl) FindAllByUser(ID uint64) (transactions []model.Transaction, err error) {
	result := t.DB.Find(&transactions).Where("id_user = ?", ID)
	return transactions, result.Error
}

// Update transaction in database
func (t *TransactionRepositoryImpl) Update(transaction *model.Transaction) error {
	var oldTransaction model.Transaction
	t.DB.Where("id = ?", transaction.ID).Find(&oldTransaction)
	oldTransaction.Description = transaction.Description
	oldTransaction.IDCategory = transaction.IDCategory
	oldTransaction.Value = transaction.Value
	result := t.DB.Save(&oldTransaction)
	return result.Error
}
