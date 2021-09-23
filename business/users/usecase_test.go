package users_test

import (
	"context"
	"kampus_merdeka/business/users"
	_mockUserRepository "kampus_merdeka/business/users/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository _mockUserRepository.Repository

var userService users.Usecase
var userDomain users.Domain

func setup() {
	userService = users.NewUserUsecase(&userRepository, time.Hour*1)
	userDomain = users.Domain{
		Id:       1,
		Name:     "Alterra",
		Email:    "alterra@gmail.com",
		Password: "abc123",
		Address:  "Malang",
		Token:    "123",
	}
}

func TestLogin(t *testing.T) {
	// t.Error("4 + 5 Harusnya 9")
	// assert.Equal(t, 2, 1)
	setup()
	t.Run("Test Case 1 | Valid Login", func(t *testing.T) {
		userRepository.On("Login",
			mock.Anything,
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(userDomain, nil).Once()

		user, err := userService.Login(context.Background(), users.Domain{
			Email:    "alterra@gmail.com",
			Password: "123",
		})
		assert.Nil(t, err)
		assert.Equal(t, "Alterra", user.Name)
	})

	t.Run("Test Case 2 | Invalid Email Empty", func(t *testing.T) {
		userRepository.On("Login",
			mock.Anything,
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(userDomain, nil).Once()

		_, err := userService.Login(context.Background(), users.Domain{
			Email:    "",
			Password: "123",
		})
		assert.NotNil(t, err)
	})

	t.Run("Test Case 3 | Invalid Password Empty", func(t *testing.T) {
		userRepository.On("Login",
			mock.Anything,
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(userDomain, nil).Once()

		_, err := userService.Login(context.Background(), users.Domain{
			Email:    "alterra@gmail.com",
			Password: "",
		})
		assert.NotNil(t, err)
	})
}
