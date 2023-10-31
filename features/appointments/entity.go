package appointments

import (
	"github.com/labstack/echo/v4"
)

type Appointment struct {
	ID              uint
	UserID          uint
	AppointmentDate string
	AppointmentTime string
}

type AppointmentInfo struct {
	UserName        string `json:"user_name"`
	AppointmentDate string `json:"appointment_date"`
	AppointmentTime string `json:"appointment_time"`
}

type UpdateAppointment struct {
	AppointmentDate string
	AppointmentTime string
}

type AppointmentHandlerInterface interface {
	GetAppointments() echo.HandlerFunc
	GetAppointment() echo.HandlerFunc
	CreateAppointment() echo.HandlerFunc
	UpdateAppointment() echo.HandlerFunc
	DeleteAppointment() echo.HandlerFunc
}

type AppointmentServiceInterface interface {
	GetAppointments(page int) ([]AppointmentInfo, int, int, error)
	GetAppointment(id int) ([]AppointmentInfo, error)
	CreateAppointment(newData Appointment) (*Appointment, error)
	UpdateAppointment(newData UpdateAppointment, id int) (bool, error)
	DeleteAppointment(id int) (bool, error)
}

type AppointmentDataInterface interface {
	GetAll(page int) ([]AppointmentInfo, int, int, error)
	GetByID(id int) ([]AppointmentInfo, error)
	Insert(newData Appointment) (*Appointment, error)
	Update(newData UpdateAppointment, id int) (bool, error)
	Delete(id int) (bool, error)
}
