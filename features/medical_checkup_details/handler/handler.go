package handler

import (
	medicalcheckupdetails "MiniProject/features/medical_checkup_details"
	"MiniProject/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MedicalCheckupDetailHandler struct {
	s medicalcheckupdetails.MedicalCheckupDetailServiceInterface
}

func NewHandler(service medicalcheckupdetails.MedicalCheckupDetailServiceInterface) medicalcheckupdetails.MedicalCheckupDetailHandlerInterface {
	return &MedicalCheckupDetailHandler{
		s: service,
	}
}

func (mcdh *MedicalCheckupDetailHandler) GetMedicalCheckupDetails() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramMCUID = c.Param("idmcu")
		idMcu, err := strconv.Atoi(paramMCUID)
		if err != nil {
			c.Logger().Fatal("Handler : Param ID MCU Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		result, err := mcdh.s.GetMedicalCheckupDetails(idMcu)

		if err != nil {
			c.Logger().Fatal("Handler : Get All Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}
func (mcdh *MedicalCheckupDetailHandler) GetMedicalCheckupDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramMCUID = c.Param("idmcu")
		idMcu, err := strconv.Atoi(paramMCUID)
		if err != nil {
			c.Logger().Fatal("Handler : Param ID MCU Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var paramMCUDID = c.Param("idmcudetail")
		idMcuDetail, err := strconv.Atoi(paramMCUDID)
		if err != nil {
			c.Logger().Fatal("Handler : Param ID MCU Detail Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		result, err := mcdh.s.GetMedicalCheckupDetail(idMcu, idMcuDetail)

		if err != nil {
			c.Logger().Fatal("Handler : Get By ID Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}
func (mcdh *MedicalCheckupDetailHandler) CreateMedicalCheckupDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramMCUID = c.Param("idmcu")
		idMcu, err := strconv.Atoi(paramMCUID)
		if err != nil {
			c.Logger().Fatal("Handler : Param ID MCU Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var input = new(InputRequest)
		if err := c.Bind(&input); err != nil {
			c.Logger().Fatal("Handler : Bind Input Error  : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var serviceInput = new(medicalcheckupdetails.MedicalCheckupDetail)
		serviceInput.MedicalCheckupID = uint(idMcu)
		serviceInput.MedicineID = input.MedicineID
		serviceInput.Quantity = input.Quantity

		result, err := mcdh.s.CreateMedicalCheckupDetail(*serviceInput)

		if err != nil {
			c.Logger().Fatal("Handler : Input Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		var response = new(InputResponse)
		response.MedicalCheckupID = result.MedicalCheckupID
		response.MedicineID = result.MedicineID
		response.Quantity = result.Quantity

		return c.JSON(http.StatusCreated, helper.FormatResponse("Success", response))
	}
}
func (mcdh *MedicalCheckupDetailHandler) UpdateMedicalCheckupDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramMCUID = c.Param("idmcu")
		idMcu, err := strconv.Atoi(paramMCUID)
		if err != nil {
			c.Logger().Fatal("Handler : Param ID MCU Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var paramMCUDID = c.Param("idmcudetail")
		idMcuDetail, err := strconv.Atoi(paramMCUDID)
		if err != nil {
			c.Logger().Fatal("Handler : Param ID MCU Detail Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var input = new(InputRequest)
		if err := c.Bind(&input); err != nil {
			c.Logger().Fatal("Handler : Bind Input Error  : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var serviceInput = new(medicalcheckupdetails.UpdateMedicalCheckupDetail)
		serviceInput.MedicineID = input.MedicineID
		serviceInput.Quantity = input.Quantity

		result, err := mcdh.s.UpdateMedicalCheckupDetail(*serviceInput, idMcu, idMcuDetail)

		if err != nil {
			c.Logger().Fatal("Handler : Update Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}
func (mcdh *MedicalCheckupDetailHandler) DeleteMedicalCheckupDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramMCUID = c.Param("idmcu")
		idMcu, err := strconv.Atoi(paramMCUID)
		if err != nil {
			c.Logger().Fatal("Handler : Param ID MCU Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var paramMCUDID = c.Param("idmcudetail")
		idMcuDetail, err := strconv.Atoi(paramMCUDID)
		if err != nil {
			c.Logger().Fatal("Handler : Param ID MCU Detail Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		result, err := mcdh.s.DeleteMedicalCheckupDetail(idMcu, idMcuDetail)

		if err != nil {
			c.Logger().Fatal("Handler : Delete Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusNoContent, helper.FormatResponse("Success", result))
	}
}
