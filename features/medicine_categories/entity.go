package medicinecategories

import "github.com/labstack/echo/v4"

type MedicineCategory struct {
	ID          uint
	Name        string
	Description string
}

type UpdateMedicineCategory struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MedicineCategoryHandlerInterface interface {
	GetMedicineCategories() echo.HandlerFunc
	GetMedicineCategory() echo.HandlerFunc
	CreateMedicineCategory() echo.HandlerFunc
	UpdateMedicineCategory() echo.HandlerFunc
	DeleteMedicineCategory() echo.HandlerFunc
}

type MedicineCategoryServiceInterface interface {
	GetMedicineCategories() ([]MedicineCategory, error)
	GetMedicineCategory(id int) ([]MedicineCategory, error)
	CreateMedicineCategory(newData MedicineCategory) (*MedicineCategory, error)
	UpdateMedicineCategory(newData UpdateMedicineCategory, id int) (bool, error)
	DeleteMedicineCategory(id int) (bool, error)
}

type MedicineCategoryDataInterface interface {
	GetAll() ([]MedicineCategory, error)
	GetByID(id int) ([]MedicineCategory, error)
	Insert(newData MedicineCategory) (*MedicineCategory, error)
	Update(newData UpdateMedicineCategory, id int) (bool, error)
	Delete(id int) (bool, error)
}
