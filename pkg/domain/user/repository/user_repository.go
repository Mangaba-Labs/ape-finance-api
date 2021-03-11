package repository

import (
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/user/model"
	"gorm.io/gorm"
)

// UserRepository Contract
type UserRepository interface {
	FindAll() (users []model.User, err error)
	FindOneByEmail(email string) (user model.User, err error)
	FindById(id int) (user model.User, err error)
	Create(user *model.User) error
	Delete(id int) error
}

//NewUserRepository repository postgres implementation
func NewUserRepository(db *gorm.DB) UserRepository {
	return &Repository{
		DB: db,
	}
}
