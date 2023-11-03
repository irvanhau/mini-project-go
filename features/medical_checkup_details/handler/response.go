package handler

type InputResponse struct {
	MedicalCheckupID uint `json:"medical_checkup_id" form:"medical_checkup_id"`
	MedicineID       uint `json:"medicine_id" form:"medicine_id"`
	Quantity         int  `json:"quantity" form:"quantity"`
}

type MedicalCheckupResponseDetail struct {
	MedicalCheckupID uint   `json:"medical_checkup_id"`
	Complain         string `json:"complain"`
	Treatment        string `json:"treatment"`
	DetailInfo       []struct {
		MedicineName string `json:"medicine_name"`
		Quantity     int    `json:"quantity"`
	} `json:"detail"`
}
