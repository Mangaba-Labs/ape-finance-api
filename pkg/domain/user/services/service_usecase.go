package services

import (
	"errors"
	"fmt"

	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/user/model"

	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/user/repository"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Repository repository.UserRepository
}

// CreateUser on app
func (s Service) CreateUser(usr *model.User) (*model.Response, error) {
	_, err := s.Repository.FindOneByEmail(usr.Email)

	// err == nil means that we find an usr with this e-mail on db
	if err == nil {
		return nil, errors.New("usr already exists")
	}

	hash, err := hashPassword(usr.Password)

	if err != nil {
		return nil, err
	}

	usr.Password = hash

	if err := s.Repository.Create(usr); err != nil {
		return nil, err
	}

	return &model.Response{
		Email: usr.Email,
		Name:  usr.Name,
	}, nil
}

func (s Service) UpdateUser(user *model.User, id int) (err error) {
	result, err := s.Repository.FindOneByEmail(user.Email)

	if err != nil {
		errMessage := fmt.Sprintf("cannot find user %s on database", user.Email)
		return errors.New(errMessage)
	}

	result.Name = user.Name
	result.Email = user.Email

	err = s.Repository.Create(&result)
	return
}

func (s Service) UpdateUserPassword(user *model.User, id int) (err error) {
	result, err := s.Repository.FindOneByEmail(user.Email)

	if err != nil {
		errMessage := fmt.Sprintf("cannot find user %s on database", user.Email)
		return errors.New(errMessage)
	}

	result.Password = user.Password

	err = s.Repository.Create(&result)
	return
}

func (s Service) DeleteUser(id int) error {
	err := s.Repository.Delete(id)

	if err != nil {
		return errors.New("user does not exist")
	}

	return nil
}

func (s Service) GetUserByEmail(email string) (usr model.User, err error) {
	usr, err = s.Repository.FindOneByEmail(email)

	if err != nil {
		errMessage := fmt.Sprintf("cannot find usr %s on database", email)
		return model.User{}, errors.New(errMessage)
	}

	return
}

func (s Service) GetUserById(id int) (usr model.User, err error) {
	usr, err = s.Repository.FindById(id)

	if err != nil {
		errMessage := fmt.Sprintf("cannot find usr %d on database", id)
		return model.User{}, errors.New(errMessage)
	}

	return
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}
