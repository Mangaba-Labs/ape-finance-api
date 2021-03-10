package services

import (
	"errors"
	"fmt"

	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/user"
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/user/repository"
	"golang.org/x/crypto/bcrypt"
)

// Service struct
type Service struct {
	Repository repository.UserRepository
}

// CreateUser on app
func (s Service) CreateUser(usr *user.User) (*user.Response, error) {
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

	return &user.Response{
		Email: usr.Email,
		Name:  usr.Name,
	}, nil
}

// UpdateUser service
func (s Service) UpdateUser(user *user.User, id int) (err error) {
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

// UpdateUserPassword changes user password
func (s Service) UpdateUserPassword(user *user.User, id int) (err error) {
	result, err := s.Repository.FindOneByEmail(user.Email)

	if err != nil {
		errMessage := fmt.Sprintf("cannot find user %s on database", user.Email)
		return errors.New(errMessage)
	}

	result.Password = user.Password

	err = s.Repository.Create(&result)
	return
}

// DeleteUser from application
func (s Service) DeleteUser(id int) error {
	err := s.Repository.Delete(id)

	if err != nil {
		return errors.New("user does not exist")
	}

	return nil
}

// GetUserByEmail find the user using his email
func (s Service) GetUserByEmail(email string) (usr user.User, err error) {
	usr, err = s.Repository.FindOneByEmail(email)

	if err != nil {
		errMessage := fmt.Sprintf("cannot find usr %s on database", email)
		return user.User{}, errors.New(errMessage)
	}

	return
}

// GetUserByID find the user using his ID
func (s Service) GetUserByID(id int) (usr user.User, err error) {
	usr, err = s.Repository.FindByID(id)

	if err != nil {
		errMessage := fmt.Sprintf("cannot find usr %d on database", id)
		return user.User{}, errors.New(errMessage)
	}

	return
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}
