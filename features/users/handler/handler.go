package handler

import (
	"MiniProject/features/users"
	"MiniProject/helper"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	s   users.UserServiceInterface
	jwt helper.JWTInterface
}

func NewHandler(service users.UserServiceInterface, j helper.JWTInterface) users.UserHandlerInterface {
	return &UserHandler{
		s:   service,
		jwt: j,
	}
}

func (uh *UserHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(RegisterRequest)
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid user input", nil))
		}

		var serviceInput = new(users.User)
		serviceInput.FullName = input.FullName
		serviceInput.Address = input.Address
		serviceInput.Email = input.Email
		serviceInput.Password = input.Password
		serviceInput.BOD = input.BOD
		serviceInput.IdentityNumber = input.IdentityNumber
		serviceInput.Role = input.Role

		result, err := uh.s.Register(*serviceInput)

		if err != nil {
			if strings.Contains(err.Error(), "Email") {
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Email has already registered", nil))
			}
			if strings.Contains(err.Error(), "Identity Number") {
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Identity Number has already registered", nil))
			}
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data", nil))
		}

		var response = new(RegisterResponse)
		response.FullName = result.FullName
		response.Email = result.Email
		response.Address = result.Address
		response.IdentityNumber = result.IdentityNumber
		response.BOD = result.BOD
		response.Role = result.Role

		return c.JSON(http.StatusCreated, helper.FormatResponse("Success Register", response))
	}
}
func (uh *UserHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(LoginInput)
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid user input", nil))
		}

		result, err := uh.s.Login(input.Email, input.Password)

		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, helper.FormatResponse("Data not found", nil))
			}
			if strings.Contains(err.Error(), "Incorrect Password") {
				return c.JSON(http.StatusNotFound, helper.FormatResponse("Password Incorrect", nil))
			}
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data", nil))
		}

		var response = new(LoginResponse)
		response.FullName = result.FullName
		response.Address = result.Address
		response.BOD = result.BOD
		response.IdentityNumber = result.IdentityNumber
		response.Email = result.Email
		response.Token = result.Access

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Login", response))
	}
}

func (uh *UserHandler) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := uh.jwt.CheckRole(c)
		fmt.Println(role)
		if role == "Patient" {
			userId := uh.jwt.GetID(c)
			userIdInt := int(userId.(float64))

			res, err := uh.s.GetUser(userIdInt)

			if err != nil {
				return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot Process Data", nil))
			}

			response := new(UserResponse)
			response.Email = res.Email
			response.FullName = res.FullName
			response.Address = res.Address
			response.BOD = res.BOD
			response.IdentityNumber = res.IdentityNumber

			return c.JSON(http.StatusOK, helper.FormatResponse("Success get users", response))
		}

		res, err := uh.s.GetUsers()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot Process Data", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success get users", res))
	}
}
