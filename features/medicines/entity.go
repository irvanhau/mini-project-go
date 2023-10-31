package medicines

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Medicine struct {
	ID           uint   `json:"id"`
	CategoryID   uint   `json:"category_id"`
	Name         string `json:"name"`
	StockMinimum int    `json:"stock_minimum"`
	Stock        int    `json:"stock"`
	Price        int    `json:"price"`
	Photo        string `json:"photo"`
	File         string `json:"file"`
}

type MedicineInfo struct {
	ID           uint   `json:"id"`
	CategoryName string `json:"category_name"`
	Name         string `json:"name"`
	StockMinimum int    `json:"stock_minimum"`
	Stock        int    `json:"stock"`
	Price        int    `json:"price"`
	Photo        string `json:"photo"`
	File         string `json:"file"`
}

type UpdateMedicine struct {
	CategoryID   uint   `json:"category_id"`
	Name         string `json:"name"`
	StockMinimum int    `json:"stock_minimum"`
	Stock        int    `json:"stock"`
	Price        int    `json:"price"`
}

type MedicineFile struct {
	File multipart.File `json:"file"`
}

type MedicinePhoto struct {
	Photo multipart.File `json:"photo"`
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
	GetMedicines(kategori int, name string) ([]MedicineInfo, error)
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
	GetAll(kategori int, name string) ([]MedicineInfo, error)
	GetByID(id int) ([]MedicineInfo, error)
	Insert(newData Medicine) (*Medicine, error)
	Update(newData UpdateMedicine, id int) (bool, error)
	Delete(id int) (bool, error)
	UpdateFileMedicine(file string, id int) (bool, error)
	UpdatePhotoMedicine(photo string, id int) (bool, error)
}
