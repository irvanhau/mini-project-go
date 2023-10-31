package handler

type InputResponse struct {
	UserID          uint   `json:"user_id" form:"user_id"`
	AppointmentDate string `json:"appointment_date" form:"appointment_date"`
	AppointmentTime string `json:"appointment_time" form:"appointment_time"`
}

type UpdateResponse struct {
	AppointmentDate string `json:"appointment_date" form:"appointment_date"`
	AppointmentTime string `json:"appointment_time" form:"appointment_time"`
}
