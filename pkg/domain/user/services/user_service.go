package services

import (
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/user"
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/user/repository"
)

type UserService interface {
	CreateUser(user *user.User) (*user.Response, error)
	UpdateUser(user *user.User, id int) error
	DeleteUser(id int) error
	GetUserByEmail(email string) (user user.User, err error)
	GetUserById(id int) (user user.User, err error)
}

func NewUserService() (service UserService) {
	r := repository.NewUserRepository()
	service = &Service{
		Repository: r,
	}
	return
}
