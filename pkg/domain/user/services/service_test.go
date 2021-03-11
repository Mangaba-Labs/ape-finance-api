package services

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/user/model"

	"github.com/Mangaba-Labs/ape-finance-api/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// TestService layer
func TestService(t *testing.T) {
	t.Run("get an valid user by e-mail", func(t *testing.T) {
		expectedResult := model.User{
			Email:    "matheus.cumpian@hotmail.com",
			Name:     "Matheus Cumpian",
			Password: "20012000",
		}

		mockedRepository := &mocks.UserRepository{}

		mockedRepository.On("FindOneByEmail", mock.Anything).Return(model.User{
			Model:    gorm.Model{},
			Email:    "matheus.cumpian@hotmail.com",
			Name:     "Matheus Cumpian",
			Password: "20012000",
		}, nil)

		userService := Service{
			Repository: mockedRepository,
		}

		result, err := userService.GetUserByEmail("matheus.cumpian@gmail.com")

		assert.Nil(t, err, "findByEmail should not throw error")
		assert.Equal(t, result, expectedResult, "wrong result")
	})

	t.Run("get an invalid user by e-mail", func(t *testing.T) {
		expectedResult := model.User{}

		mockedRepository := &mocks.UserRepository{}

		mockedRepository.On("FindOneByEmail", mock.Anything).Return(model.User{}, errors.New("user does not exists"))

		userService := Service{
			Repository: mockedRepository,
		}

		fakeEmail := "matheus.cumpian@hotmail.com"

		result, err := userService.GetUserByEmail(fakeEmail)

		assert.Equal(t, result, expectedResult, "user must be empty")
		assert.NotNil(t, err, "must return error")
		assert.Equal(t, err.Error(), fmt.Sprintf("cannot find usr %s on database", fakeEmail))
	})
}
