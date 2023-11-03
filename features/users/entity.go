package users

import "github.com/labstack/echo/v4"

type User struct {
	ID             uint   `json:"id"`
	Email          string `json:"email"`
	IdentityNumber string `json:"identity_number"`
	FullName       string `json:"full_name"`
	Password       string `json:"password"`
	BOD            string `json:"bod"`
	Address        string `json:"address"`
	Role           string `json:"role"`
}

type UserCredential struct {
	Email          string         `json:"email"`
	IdentityNumber string         `json:"identity_number"`
	FullName       string         `json:"full_name"`
	BOD            string         `json:"bod"`
	Address        string         `json:"address"`
	Access         map[string]any `json:"access"`
}

type UserInfo struct {
	Email          string `json:"email"`
	IdentityNumber string `json:"identity_number"`
	FullName       string `json:"full_name"`
	BOD            string `json:"bod"`
	Address        string `json:"address"`
}

type UserHandlerInterface interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	GetUser() echo.HandlerFunc
}

type UserServiceInterface interface {
	Register(newData User) (*User, error)
	Login(email, password string) (*UserCredential, error)
	GetUsers() ([]UserInfo, error)
	GetUser(idUser int) (User, error)
}

type UserDataInterface interface {
	Register(newData User) (*User, error)
	Login(email, password string) (*User, error)
	GetUsers() ([]UserInfo, error)
	GetUser(idUser int) (User, error)
}
