package repository

import (
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/transaction/model"
	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
	DB *gorm.DB
}

func (i *TransactionRepositoryImpl) Create(transaction *model.Transaction) error {
	result := i.DB.Create(&transaction)
	return result.Error
}

func (i *TransactionRepositoryImpl) Delete(ID uint) error {
	result := i.DB.Where("id = ?", ID).Delete(&model.Transaction{})
	return result.Error
}

func (i *TransactionRepositoryImpl) FindByID(ID uint) (model.Transaction, error) {
	var transaction model.Transaction
	result := i.DB.Find(&transaction).Where("id = ?", ID)
	return transaction, result.Error
}

func (i *TransactionRepositoryImpl) FindAllByUser(ID uint64) (transactions []model.Transaction, err error) {
	result := i.DB.Find(&transactions).Where("id_user = ?", ID)
	return transactions, result.Error
}

func (i *TransactionRepositoryImpl) Update(transaction *model.Transaction) error {
	return nil
}
