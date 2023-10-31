package service

import (
	"MiniProject/features/appointments"
	"errors"
)

type AppointmentService struct {
	data appointments.AppointmentDataInterface
}

func New(data appointments.AppointmentDataInterface) appointments.AppointmentServiceInterface {
	return &AppointmentService{
		data: data,
	}
}

func (as *AppointmentService) GetAppointments(page int) ([]appointments.AppointmentInfo, int, int, error) {
	result, totalData, totalPage, err := as.data.GetAll(page)

	if err != nil {
		return nil, 0, 0, errors.New("Get All Failed")
	}

	return result, totalData, totalPage, nil
}
func (as *AppointmentService) GetAppointment(id int) ([]appointments.AppointmentInfo, error) {
	result, err := as.data.GetByID(id)

	if err != nil {
		return nil, errors.New("Get By ID Failed")
	}

	return result, nil
}
func (as *AppointmentService) CreateAppointment(newData appointments.Appointment) (*appointments.Appointment, error) {

	result, err := as.data.Insert(newData)

	if err != nil {
		return nil, errors.New("Insert Process Failed")
	}

	return result, nil
}
func (as *AppointmentService) UpdateAppointment(newData appointments.UpdateAppointment, id int) (bool, error) {
	result, err := as.data.Update(newData, id)

	if err != nil {
		return false, errors.New("Update Process Failed")
	}

	return result, nil
}
func (as *AppointmentService) DeleteAppointment(id int) (bool, error) {
	result, err := as.data.Delete(id)

	if err != nil {
		return false, errors.New("Delete Process Failed")
	}

	return result, nil
}
