package handler

import (
	"MiniProject/features/users"
	"MiniProject/helper"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	s users.UserServiceInterface
}

func NewHandler(service users.UserServiceInterface) users.UserHandlerInterface {
	return &UserHandler{
		s: service,
	}
}

func (uh *UserHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(RegisterRequest)
		if err := c.Bind(&input); err != nil {
			c.Logger().Fatal("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
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
			c.Logger().Fatal("Handler : Input Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		var response = new(RegisterResponse)
		response.FullName = result.FullName
		response.Email = result.Email
		response.Address = result.Address
		response.IdentityNumber = result.IdentityNumber
		response.BOD = result.BOD
		response.Role = result.Role

		return c.JSON(http.StatusCreated, helper.FormatResponse("Success", response))
	}
}
func (uh *UserHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(LoginInput)
		if err := c.Bind(&input); err != nil {
			c.Logger().Fatal("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		result, err := uh.s.Login(input.Email, input.Password)

		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, helper.FormatResponse("Fail", nil))
			}
			if strings.Contains(err.Error(), "Incorrect Password") {
				return c.JSON(http.StatusNotFound, helper.FormatResponse("Password Incorrect", nil))
			}
			c.Logger().Fatal("Handler : Login Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		var response = new(LoginResponse)
		response.FullName = result.FullName
		response.Address = result.Address
		response.BOD = result.BOD
		response.IdentityNumber = result.IdentityNumber
		response.Email = result.Email
		response.Token = result.Access

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", response))
	}
}
