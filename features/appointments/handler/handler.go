package handler

import (
	"MiniProject/features/appointments"
	"MiniProject/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AppointmentHandler struct {
	s appointments.AppointmentServiceInterface
}

func NewHandler(service appointments.AppointmentServiceInterface) appointments.AppointmentHandlerInterface {
	return &AppointmentHandler{
		s: service,
	}
}

func (ah *AppointmentHandler) GetAppointments() echo.HandlerFunc {
	return func(c echo.Context) error {
		paramPage := c.QueryParam("page")
		page, _ := strconv.Atoi(paramPage)

		result, totalData, totalPage, err := ah.s.GetAppointments(page)

		if err != nil {
			c.Logger().Fatal("Handler : Get All Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		if page != 0 {
			return c.JSON(http.StatusOK, helper.FormatPagination("Success", result, page, totalPage, totalData))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))

	}
}

func (ah *AppointmentHandler) GetAppointment() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)

		if err != nil {
			c.Logger().Fatal("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		result, err := ah.s.GetAppointment(id)

		if err != nil {
			c.Logger().Fatal("Handler : Get By ID Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}

func (ah *AppointmentHandler) CreateAppointment() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(InputRequest)
		if err := c.Bind(&input); err != nil {
			c.Logger().Fatal("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var serviceInput = new(appointments.Appointment)
		serviceInput.UserID = input.UserID
		serviceInput.AppointmentDate = input.AppointmentDate
		serviceInput.AppointmentTime = input.AppointmentTime

		result, err := ah.s.CreateAppointment(*serviceInput)

		if err != nil {
			c.Logger().Fatal("Handler : Input Process Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var response = new(InputResponse)
		response.UserID = result.UserID
		response.AppointmentDate = result.AppointmentDate
		response.AppointmentTime = result.AppointmentTime

		return c.JSON(http.StatusCreated, helper.FormatResponse("Success", response))
	}

}
func (ah *AppointmentHandler) UpdateAppointment() echo.HandlerFunc {
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

		var serviceInput = new(appointments.UpdateAppointment)
		serviceInput.AppointmentDate = input.AppointmentDate
		serviceInput.AppointmentTime = input.AppointmentTime

		result, err := ah.s.UpdateAppointment(*serviceInput, id)

		if err != nil {
			c.Logger().Fatal("Handler : Update Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}
func (ah *AppointmentHandler) DeleteAppointment() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)

		if err != nil {
			c.Logger().Fatal("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		result, err := ah.s.DeleteAppointment(id)

		if err != nil {
			c.Logger().Fatal("Handler : Delete Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusNoContent, helper.FormatResponse("Success", result))
	}
}
