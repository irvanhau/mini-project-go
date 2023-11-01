package handler

import (
	medicalcheckups "MiniProject/features/medical_checkups"
	"MiniProject/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MedicalCheckupHandler struct {
	s medicalcheckups.MedicalCheckupServiceInterface
}

func NewHandler(service medicalcheckups.MedicalCheckupServiceInterface) medicalcheckups.MedicalCheckupHandlerInterface {
	return &MedicalCheckupHandler{
		s: service,
	}
}

func (mch *MedicalCheckupHandler) GetMedicalCheckups() echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := mch.s.GetMedicalCheckups()

		if err != nil {
			c.Logger().Fatal("Handler : Get All Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}
func (mch *MedicalCheckupHandler) GetMedicalCheckup() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			c.Logger().Fatal("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		result, err := mch.s.GetMedicalCheckup(id)

		if err != nil {
			c.Logger().Fatal("Handler : Get By ID Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}
func (mch *MedicalCheckupHandler) CreateMedicalCheckup() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(InputRequest)
		if err := c.Bind(&input); err != nil {
			c.Logger().Fatal("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var serviceInput = new(medicalcheckups.MedicalCheckup)
		serviceInput.UserID = input.UserID
		serviceInput.Complain = input.Complain
		serviceInput.Treatment = input.Treatment

		result, err := mch.s.CreateMedicalCheckup(*serviceInput)

		if err != nil {
			c.Logger().Fatal("Handler : Input Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		var response = new(InputResponse)
		response.UserID = result.UserID
		response.Complain = result.Complain
		response.Treatment = result.Treatment

		return c.JSON(http.StatusCreated, helper.FormatResponse("Success", response))
	}
}
func (mch *MedicalCheckupHandler) UpdateMedicalCheckup() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			c.Logger().Fatal("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var input = new(UpdateRequest)
		if err := c.Bind(&input); err != nil {
			c.Logger().Fatal("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var updateInput = new(medicalcheckups.UpdateMedicalCheckup)
		updateInput.Complain = input.Complain
		updateInput.Treatment = input.Treatment

		result, err := mch.s.UpdateMedicalCheckup(*updateInput, id)
		if err != nil {
			c.Logger().Fatal("Handler : Update Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}
func (mch *MedicalCheckupHandler) DeleteMedicalCheckup() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			c.Logger().Fatal("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		result, err := mch.s.DeleteMedicalCheckup(id)
		if err != nil {
			c.Logger().Fatal("Handler : Delete Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusNoContent, helper.FormatResponse("Success", result))
	}
}
