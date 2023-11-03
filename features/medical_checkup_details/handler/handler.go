package handler

import (
	medicalcheckupdetails "MiniProject/features/medical_checkup_details"
	"MiniProject/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MedicalCheckupDetailHandler struct {
	s   medicalcheckupdetails.MedicalCheckupDetailServiceInterface
	jwt helper.JWTInterface
}

func NewHandler(service medicalcheckupdetails.MedicalCheckupDetailServiceInterface, jwt helper.JWTInterface) medicalcheckupdetails.MedicalCheckupDetailHandlerInterface {
	return &MedicalCheckupDetailHandler{
		s:   service,
		jwt: jwt,
	}
}

func (mcdh *MedicalCheckupDetailHandler) GetMedicalCheckupDetails() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := mcdh.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}

		var paramMCUID = c.Param("idmcu")
		idMcu, err := strconv.Atoi(paramMCUID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid param id medical checkup", nil))
		}

		result, resMed, err := mcdh.s.GetMedicalCheckupDetails(idMcu)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data", nil))
		}

		response := MedicalCheckupResponseDetail{
			MedicalCheckupID: result.MedicalCheckupID,
			Complain:         result.Complain,
			Treatment:        result.Treatment,
			DetailInfo: make([]struct {
				MedicineName string "json:\"medicine_name\""
				Quantity     int    "json:\"quantity\""
			}, len(resMed)),
		}

		for i, medicine := range resMed {
			response.DetailInfo[i].MedicineName = medicine.MedicineName
			response.DetailInfo[i].Quantity = medicine.Quantity
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success get medical checkup detail", response))
	}
}
func (mcdh *MedicalCheckupDetailHandler) GetMedicalCheckupDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := mcdh.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}

		var paramMCUID = c.Param("idmcu")
		idMcu, err := strconv.Atoi(paramMCUID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid param id medical checkup", nil))
		}

		var paramMCUDID = c.Param("idmcudetail")
		idMcuDetail, err := strconv.Atoi(paramMCUDID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid param id medical checkup detail", nil))
		}

		result, err := mcdh.s.GetMedicalCheckupDetail(idMcu, idMcuDetail)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success get medical checkup detail", result))
	}
}
func (mcdh *MedicalCheckupDetailHandler) CreateMedicalCheckupDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := mcdh.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}

		var paramMCUID = c.Param("idmcu")
		idMcu, err := strconv.Atoi(paramMCUID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid param id medical checkup", nil))
		}

		var input = new(InputRequest)
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid user input", nil))
		}

		var serviceInput = new(medicalcheckupdetails.MedicalCheckupDetail)
		serviceInput.MedicalCheckupID = uint(idMcu)
		serviceInput.MedicineID = input.MedicineID
		serviceInput.Quantity = input.Quantity

		result, err := mcdh.s.CreateMedicalCheckupDetail(*serviceInput)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data", nil))
		}

		var response = new(InputResponse)
		response.MedicalCheckupID = result.MedicalCheckupID
		response.MedicineID = result.MedicineID
		response.Quantity = result.Quantity

		return c.JSON(http.StatusCreated, helper.FormatResponse("Success created medical checkup detail", response))
	}
}
func (mcdh *MedicalCheckupDetailHandler) UpdateMedicalCheckupDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := mcdh.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}

		var paramMCUID = c.Param("idmcu")
		idMcu, err := strconv.Atoi(paramMCUID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid param id medical checkup", nil))
		}

		var paramMCUDID = c.Param("idmcudetail")
		idMcuDetail, err := strconv.Atoi(paramMCUDID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid param id medical checkup detail", nil))
		}

		var input = new(InputRequest)
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid user input", nil))
		}

		var serviceInput = new(medicalcheckupdetails.UpdateMedicalCheckupDetail)
		serviceInput.MedicineID = input.MedicineID
		serviceInput.Quantity = input.Quantity

		result, err := mcdh.s.UpdateMedicalCheckupDetail(*serviceInput, idMcu, idMcuDetail)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success update medical checkup detail", result))
	}
}
func (mcdh *MedicalCheckupDetailHandler) DeleteMedicalCheckupDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := mcdh.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}

		var paramMCUID = c.Param("idmcu")
		idMcu, err := strconv.Atoi(paramMCUID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid param id medical checkup", nil))
		}

		var paramMCUDID = c.Param("idmcudetail")
		idMcuDetail, err := strconv.Atoi(paramMCUDID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid param id medical checkup detail", nil))
		}

		result, err := mcdh.s.DeleteMedicalCheckupDetail(idMcu, idMcuDetail)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data", nil))
		}

		return c.JSON(http.StatusNoContent, helper.FormatResponse("Success delete medical checkup detail", result))
	}
}
