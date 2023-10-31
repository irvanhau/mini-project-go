package users

import "github.com/labstack/echo/v4"

type User struct {
	ID             uint
	Email          string
	Password       string
	IdentityNumber string
	FullName       string
	BOD            string
	Address        string
	Role           string
}

type UserCredential struct {
	Email          string
	IdentityNumber string
	FullName       string
	BOD            string
	Address        string
	Access         map[string]any
}

type UserHandlerInterface interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
}

type UserServiceInterface interface {
	Register(newData User) (*User, error)
	Login(email, password string) (*UserCredential, error)
}

type UserDataInterface interface {
	Register(newData User) (*User, error)
	Login(email, password string) (*User, error)
}
