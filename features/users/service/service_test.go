package service

import (
	"MiniProject/features/users"
	"MiniProject/features/users/mocks"
	helper "MiniProject/helper/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	jwt := helper.NewJWTInterface(t)
	data := mocks.NewUserDataInterface(t)
	service := New(data, jwt)
	newUser := users.User{
		FullName: "Irvan Hauwerich",
		Email:    "irvanhau@gmail.com",
		Password: "vanhau123",
		Role:     "Patient",
		Address:  "Cikarang",
		BOD:      "1999-09-25",
	}

	t.Run("Success Register", func(t *testing.T) {
		data.On("Register", newUser).Return(&newUser, nil).Once()

		result, err := service.Register(newUser)
		assert.Nil(t, err)
		assert.Equal(t, newUser.Email, result.Email)
		assert.Equal(t, newUser.FullName, result.FullName)
		assert.Equal(t, newUser.Password, result.Password)
		assert.Equal(t, newUser.Role, result.Role)
		assert.Equal(t, newUser.Address, result.Address)
		assert.Equal(t, newUser.BOD, result.BOD)
		data.AssertExpectations(t)
	})

	t.Run("Input Error", func(t *testing.T) {
		data.On("Register", newUser).Return(nil, errors.New("Register Process Failed")).Once()

		result, err := service.Register(newUser)
		assert.Error(t, err)
		assert.EqualError(t, err, "Register Process Failed")
		assert.Nil(t, result)
	})

	t.Run("Email Duplicate", func(t *testing.T) {
		data.On("Register", newUser).Return(nil, errors.New("Email has already registered")).Once()

		result, err := service.Register(newUser)
		assert.Error(t, err)
		assert.EqualError(t, err, "Email has already registered")
		assert.Nil(t, result)
	})

	t.Run("Identity Number Duplicate", func(t *testing.T) {
		data.On("Register", newUser).Return(nil, errors.New("Identity Number has already registered")).Once()

		result, err := service.Register(newUser)
		assert.Error(t, err)
		assert.EqualError(t, err, "Identity Number has already registered")
		assert.Nil(t, result)
	})
}

func TestLogin(t *testing.T) {
	jwt := helper.NewJWTInterface(t)
	data := mocks.NewUserDataInterface(t)
	service := New(data, jwt)
	userData := users.User{
		FullName: "Irvan Hauwerich",
		Email:    "irvanhau@gmail.com",
		Password: "vanhau123",
	}

	t.Run("Success Login", func(t *testing.T) {
		jwtResult := map[string]any{"access_token": uint(0), "role": "mockToken"}
		data.On("Login", userData.Email, userData.Password).Return(&userData, nil).Once()
		jwt.On("GenerateJWT", uint(0), "").Return(jwtResult).Once()
		result, err := service.Login(userData.Email, userData.Password)

		data.AssertExpectations(t)
		jwt.AssertExpectations(t)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, userData.Email, result.Email)
		assert.Equal(t, "Irvan Hauwerich", result.FullName)
		assert.Equal(t, jwtResult, result.Access)
	})

	t.Run("Input Error", func(t *testing.T) {
		userFail := users.User{
			Email:    "irvanhau",
			Password: "vanhau123",
		}

		data.On("Login", userFail.Email, userFail.Password).Return(nil, errors.New("data not found")).Once()

		result, err := service.Login(userFail.Email, userFail.Password)

		assert.Error(t, err)
		assert.EqualError(t, err, "data not found")
		assert.Nil(t, result)
	})

	t.Run("Password Incorrect", func(t *testing.T) {
		userFail := users.User{
			Email:    "irvanhau",
			Password: "vanhau123",
		}

		data.On("Login", userFail.Email, userFail.Password).Return(nil, errors.New("Incorrect Password")).Once()

		result, err := service.Login(userFail.Email, userFail.Password)

		assert.Error(t, err)
		assert.EqualError(t, err, "Incorrect Password")
		assert.Nil(t, result)
	})

	t.Run("Server Error", func(t *testing.T) {
		userFail := users.User{
			Email:    "irvanhau",
			Password: "vanhau123",
		}

		data.On("Login", userFail.Email, userFail.Password).Return(nil, errors.New("Process Failed")).Once()

		result, err := service.Login(userFail.Email, userFail.Password)

		assert.Error(t, err)
		assert.EqualError(t, err, "Process Failed")
		assert.Nil(t, result)
	})

	t.Run("Token Failed", func(t *testing.T) {
		data.On("Login", userData.Email, userData.Password).Return(&userData, nil).Once()
		jwt.On("GenerateJWT", uint(0), "").Return(nil).Once()
		result, err := service.Login(userData.Email, userData.Password)

		assert.Error(t, err)
		assert.EqualError(t, err, "Token Process Failed")
		assert.Nil(t, result)
	})
}

func TestGetUsers(t *testing.T) {
	data := mocks.NewUserDataInterface(t)
	jwt := helper.NewJWTInterface(t)
	service := New(data, jwt)
	user := []users.UserInfo{}

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetUsers").Return(user, nil).Once()

		result, err := service.GetUsers()
		assert.Nil(t, err)
		assert.NotNil(t, result)
		data.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetUsers").Return(nil, errors.New("Get Users Failed")).Once()

		result, err := service.GetUsers()
		assert.Error(t, err)
		assert.EqualError(t, err, "Get Users Failed")
		assert.Nil(t, result)
	})
}

func TestGetUser(t *testing.T) {
	data := mocks.NewUserDataInterface(t)
	jwt := helper.NewJWTInterface(t)
	service := New(data, jwt)
	user := users.User{}

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetUser", 1).Return(user, nil).Once()

		result, err := service.GetUser(1)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		data.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetUser", 1).Return(user, errors.New("Get User By ID JWT Failed")).Once()

		result, err := service.GetUser(1)
		assert.Error(t, err)
		assert.EqualError(t, err, "Get User By ID JWT Failed")
		assert.NotNil(t, result)
	})
}
