package handler

type InputRequest struct {
	MedicineID uint `json:"medicine_id" form:"medicine_id"`
	Quantity   int  `json:"quantity" form:"quantity"`
}
