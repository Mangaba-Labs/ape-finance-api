package repository

import (
	"github.com/Mangaba-Labs/tempoo-api/pkg/domain/database"
	"github.com/Mangaba-Labs/tempoo-api/pkg/domain/user"
	"gorm.io/gorm"
)

// UserRepository Contract
type UserRepository interface {
	FindAll() (users *gorm.DB, err error)
	FindOneByEmail(email string) (user user.User, err error)
	FindById(id int) (user user.User, err error)
	Create(user *user.User) error
	Delete(id int) error
}

//NewUserRepository repository postgres implementation
func NewUserRepository() UserRepository {
	return &Repository{
		DB: database.Instance,
	}
}
