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
		page, err := strconv.Atoi(paramPage)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid user input", nil))
		}

		result, totalData, totalPage, err := ah.s.GetAppointments(page)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data", nil))
		}

		if page != 0 {
			return c.JSON(http.StatusOK, helper.FormatPagination("Success get appointment", result, page, totalPage, totalData))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success get appointment", result))

	}
}

func (ah *AppointmentHandler) GetAppointment() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid user input", nil))
		}

		result, err := ah.s.GetAppointment(id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success get appointment", result))
	}
}

func (ah *AppointmentHandler) CreateAppointment() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(InputRequest)
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid user input", nil))
		}

		var serviceInput = new(appointments.Appointment)
		serviceInput.UserID = input.UserID
		serviceInput.AppointmentDate = input.AppointmentDate
		serviceInput.AppointmentTime = input.AppointmentTime

		result, err := ah.s.CreateAppointment(*serviceInput)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Cannot process data", nil))
		}

		var response = new(InputResponse)
		response.UserID = result.UserID
		response.AppointmentDate = result.AppointmentDate
		response.AppointmentTime = result.AppointmentTime

		return c.JSON(http.StatusCreated, helper.FormatResponse("Success created appointment", response))
	}

}
func (ah *AppointmentHandler) UpdateAppointment() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid user input", nil))
		}

		var input = new(UpdateRequest)
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid user input", nil))
		}

		var serviceInput = new(appointments.UpdateAppointment)
		serviceInput.AppointmentDate = input.AppointmentDate
		serviceInput.AppointmentTime = input.AppointmentTime

		result, err := ah.s.UpdateAppointment(*serviceInput, id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success update appointment", result))
	}
}
func (ah *AppointmentHandler) DeleteAppointment() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid user input", nil))
		}

		result, err := ah.s.DeleteAppointment(id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data", nil))
		}

		return c.JSON(http.StatusNoContent, helper.FormatResponse("Success delete appointment", result))
	}
}
