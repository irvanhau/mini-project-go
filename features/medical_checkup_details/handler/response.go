package handler

type InputResponse struct {
	MedicalCheckupID uint `json:"medical_checkup_id" form:"medical_checkup_id"`
	MedicineID       uint `json:"medicine_id" form:"medicine_id"`
	Quantity         int  `json:"quantity" form:"quantity"`
}
