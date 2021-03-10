package services

import (
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/user"
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/user/repository"
)

// UserService interface implementation
type UserService interface {
	CreateUser(user *user.User) (*user.Response, error)
	UpdateUser(user *user.User, id int) error
	DeleteUser(id int) error
	GetUserByEmail(email string) (user user.User, err error)
	GetUserByID(id int) (user user.User, err error)
}

// NewUserService create new instance of user service
func NewUserService() (service UserService) {
	r := repository.NewUserRepository()
	service = &Service{
		Repository: r,
	}
	return
}
