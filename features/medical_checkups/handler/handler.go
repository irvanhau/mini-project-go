package handler

import (
	medicalcheckups "MiniProject/features/medical_checkups"
	"MiniProject/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MedicalCheckupHandler struct {
	s   medicalcheckups.MedicalCheckupServiceInterface
	jwt helper.JWTInterface
}

func NewHandler(service medicalcheckups.MedicalCheckupServiceInterface, jwt helper.JWTInterface) medicalcheckups.MedicalCheckupHandlerInterface {
	return &MedicalCheckupHandler{
		s:   service,
		jwt: jwt,
	}
}

func (mch *MedicalCheckupHandler) GetMedicalCheckups() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := mch.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}

		result, err := mch.s.GetMedicalCheckups()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success get medical checkup", result))
	}
}
func (mch *MedicalCheckupHandler) GetMedicalCheckup() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := mch.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}

		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid param id medical checkup", nil))
		}

		result, err := mch.s.GetMedicalCheckup(id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success get medical checkup", result))
	}
}
func (mch *MedicalCheckupHandler) CreateMedicalCheckup() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := mch.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}

		var input = new(InputRequest)
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid user input", nil))
		}

		var serviceInput = new(medicalcheckups.MedicalCheckup)
		serviceInput.UserID = input.UserID
		serviceInput.Complain = input.Complain
		serviceInput.Treatment = input.Treatment

		result, err := mch.s.CreateMedicalCheckup(*serviceInput)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data", nil))
		}

		var response = new(InputResponse)
		response.UserID = result.UserID
		response.Complain = result.Complain
		response.Treatment = result.Treatment

		return c.JSON(http.StatusCreated, helper.FormatResponse("Success created medical checkup", response))
	}
}
func (mch *MedicalCheckupHandler) UpdateMedicalCheckup() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := mch.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}

		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid param id medical checkup", nil))
		}

		var input = new(UpdateRequest)
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid user input", nil))
		}

		var updateInput = new(medicalcheckups.UpdateMedicalCheckup)
		updateInput.Complain = input.Complain
		updateInput.Treatment = input.Treatment

		result, err := mch.s.UpdateMedicalCheckup(*updateInput, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success update medical checkup", result))
	}
}
func (mch *MedicalCheckupHandler) DeleteMedicalCheckup() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := mch.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}

		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid param id medical checkup", nil))
		}

		result, err := mch.s.DeleteMedicalCheckup(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data", nil))
		}

		return c.JSON(http.StatusNoContent, helper.FormatResponse("Success delete medical checkup", result))
	}
}
