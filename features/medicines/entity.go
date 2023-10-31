package medicines

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Medicine struct {
	ID           uint
	CategoryID   uint
	Name         string
	StockMinimum int
	Stock        int
	Price        int
	Photo        string
	File         string
}

type MedicineInfo struct {
	ID           uint
	CategoryName string `json:"category_name"`
	Name         string
	StockMinimum int
	Stock        int
	Price        int
	Photo        string
	File         string
}

type UpdateMedicine struct {
	CategoryID   uint
	Name         string
	StockMinimum int
	Stock        int
	Price        int
}

type MedicineFile struct {
	File multipart.File
}

type MedicinePhoto struct {
	Photo multipart.File
}

type MedicineHandlerInterface interface {
	GetMedicines() echo.HandlerFunc
	GetMedicine() echo.HandlerFunc
	CreateMedicine() echo.HandlerFunc
	UpdateMedicine() echo.HandlerFunc
	DeleteMedicine() echo.HandlerFunc
	UpdateFileMedicine() echo.HandlerFunc
	UpdatePhotoMedicine() echo.HandlerFunc
}

type MedicineServiceInterface interface {
	GetMedicines() ([]MedicineInfo, error)
	GetMedicine(id int) ([]MedicineInfo, error)
	CreateMedicine(newData Medicine) (*Medicine, error)
	UpdateMedicine(newData UpdateMedicine, id int) (bool, error)
	DeleteMedicine(id int) (bool, error)
	FileUpload(file MedicineFile) (string, error)
	PhotoUpload(file MedicinePhoto) (string, error)
	UpdateFileMedicine(file string, id int) (bool, error)
	UpdatePhotoMedicine(photo string, id int) (bool, error)
}

type MedicineDataInterface interface {
	GetAll() ([]MedicineInfo, error)
	GetByID(id int) ([]MedicineInfo, error)
	Insert(newData Medicine) (*Medicine, error)
	Update(newData UpdateMedicine, id int) (bool, error)
	Delete(id int) (bool, error)
	UpdateFileMedicine(file string, id int) (bool, error)
	UpdatePhotoMedicine(photo string, id int) (bool, error)
}
