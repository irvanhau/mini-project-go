package medicalcheckups

import "github.com/labstack/echo/v4"

type MedicalCheckup struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	Complain  string `json:"complain"`
	Treatment string `json:"treatment"`
}

type MedicalCheckupInfo struct {
	UserName  string `json:"user_name"`
	Complain  string `json:"complain"`
	Treatment string `json:"treatment"`
}

type UpdateMedicalCheckup struct {
	Complain  string `json:"complain"`
	Treatment string `json:"treatment"`
}

type MedicalCheckupHandlerInterface interface {
	GetMedicalCheckups() echo.HandlerFunc
	GetMedicalCheckup() echo.HandlerFunc
	CreateMedicalCheckup() echo.HandlerFunc
	UpdateMedicalCheckup() echo.HandlerFunc
	DeleteMedicalCheckup() echo.HandlerFunc
}

type MedicalCheckupServiceInterface interface {
	GetMedicalCheckups() ([]MedicalCheckupInfo, error)
	GetMedicalCheckup(id int) ([]MedicalCheckupInfo, error)
	CreateMedicalCheckup(newData MedicalCheckup) (*MedicalCheckup, error)
	UpdateMedicalCheckup(newData UpdateMedicalCheckup, id int) (bool, error)
	DeleteMedicalCheckup(id int) (bool, error)
}

type MedicalCheckupDataInterface interface {
	GetAll() ([]MedicalCheckupInfo, error)
	GetByID(id int) ([]MedicalCheckupInfo, error)
	Insert(newData MedicalCheckup) (*MedicalCheckup, error)
	Update(newData UpdateMedicalCheckup, id int) (bool, error)
	Delete(id int) (bool, error)
}
