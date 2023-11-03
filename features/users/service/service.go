package service

import (
	"MiniProject/features/users"
	"MiniProject/helper"
	"errors"
	"strings"
)

type UserService struct {
	d users.UserDataInterface
	j helper.JWTInterface
}

func New(data users.UserDataInterface, jwt helper.JWTInterface) users.UserServiceInterface {
	return &UserService{
		d: data,
		j: jwt,
	}
}

func (us *UserService) Register(newData users.User) (*users.User, error) {
	result, err := us.d.Register(newData)

	if err != nil {
		if strings.Contains(err.Error(), "Email") {
			return nil, errors.New("Email has already registered")
		}
		if strings.Contains(err.Error(), "Identity Number") {
			return nil, errors.New("Identity Number has already registered")
		}
		return nil, errors.New("Register Process Failed")
	}

	return result, nil
}
func (us *UserService) Login(email, password string) (*users.UserCredential, error) {
	result, err := us.d.Login(email, password)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, errors.New("data not found")
		}
		if strings.Contains(err.Error(), "Incorrect Password") {
			return nil, errors.New("Incorrect Password")
		}
		return nil, errors.New("Process Failed")
	}

	tokenData := us.j.GenerateJWT(result.ID, result.Role)

	if tokenData == nil {
		return nil, errors.New("Token Process Failed")
	}

	response := new(users.UserCredential)
	response.Email = result.Email
	response.FullName = result.FullName
	response.Address = result.Address
	response.IdentityNumber = result.IdentityNumber
	response.Access = tokenData
	response.BOD = result.BOD

	return response, nil
}

func (us *UserService) GetUsers() ([]users.UserInfo, error) {
	result, err := us.d.GetUsers()

	if err != nil {
		return result, errors.New("Get Users Failed")
	}

	return result, nil
}
func (us *UserService) GetUser(idUser int) (users.User, error) {
	result, err := us.d.GetUser(idUser)

	if err != nil {
		return result, errors.New("Get User Failed")
	}

	return result, nil
}
