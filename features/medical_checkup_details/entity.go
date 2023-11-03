package medicalcheckupdetails

import "github.com/labstack/echo/v4"

type MedicalCheckupDetail struct {
	ID               uint `json:"id"`
	MedicalCheckupID uint `json:"medical_checkup_id"`
	MedicineID       uint `json:"medicine_id"`
	Quantity         int  `json:"quantity"`
}

type MedicalCheckupDetailInfo struct {
	MedicalCheckupID uint   `json:"medical_checkup_id"`
	Complain         string `json:"complain"`
	Treatment        string `json:"treatment"`
	MedicineName     string `json:"medicine_name"`
	Quantity         int    `json:"quantity"`
}

type DetailInfo []struct {
	MedicineName string `json:"medicine_name"`
	Quantity     int    `json:"quantity"`
}

type UpdateMedicalCheckupDetail struct {
	MedicineID uint `json:"medicine_id"`
	Quantity   int  `json:"quantity"`
}

type MedicalCheckupDetailHandlerInterface interface {
	GetMedicalCheckupDetails() echo.HandlerFunc
	GetMedicalCheckupDetail() echo.HandlerFunc
	CreateMedicalCheckupDetail() echo.HandlerFunc
	UpdateMedicalCheckupDetail() echo.HandlerFunc
	DeleteMedicalCheckupDetail() echo.HandlerFunc
}

type MedicalCheckupDetailServiceInterface interface {
	GetMedicalCheckupDetails(idMcu int) (MedicalCheckupDetailInfo, DetailInfo, error)
	GetMedicalCheckupDetail(idMcu, idMcuDetail int) ([]MedicalCheckupDetailInfo, error)
	CreateMedicalCheckupDetail(newData MedicalCheckupDetail) (*MedicalCheckupDetail, error)
	UpdateMedicalCheckupDetail(newData UpdateMedicalCheckupDetail, idMcu, idMcuDetail int) (bool, error)
	DeleteMedicalCheckupDetail(idMcu, idMcuDetail int) (bool, error)
}

type MedicalCheckupDetailDataInterface interface {
	GetAll(idMcu int) (MedicalCheckupDetailInfo, DetailInfo, error)
	GetByID(idMcu, idMcuDetail int) ([]MedicalCheckupDetailInfo, error)
	Insert(newData MedicalCheckupDetail) (*MedicalCheckupDetail, error)
	Update(newData UpdateMedicalCheckupDetail, idMcu, idMcuDetail int) (bool, error)
	Delete(idMcu, idMcuDetail int) (bool, error)
}
