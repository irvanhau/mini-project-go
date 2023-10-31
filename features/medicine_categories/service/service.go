package service

import (
	medicinecategories "MiniProject/features/medicine_categories"
	"errors"
)

type MedicineCategoryService struct {
	data medicinecategories.MedicineCategoryDataInterface
}

func New(data medicinecategories.MedicineCategoryDataInterface) medicinecategories.MedicineCategoryServiceInterface {
	return &MedicineCategoryService{
		data: data,
	}
}

func (mcs *MedicineCategoryService) GetMedicineCategories() ([]medicinecategories.MedicineCategory, error) {
	result, err := mcs.data.GetAll()

	if err != nil {
		return nil, errors.New("Get All Process Failed")
	}

	return result, nil
}
func (mcs *MedicineCategoryService) GetMedicineCategory(id int) ([]medicinecategories.MedicineCategory, error) {
	result, err := mcs.data.GetByID(id)

	if err != nil {
		return nil, errors.New("Get By ID Failed")
	}

	return result, nil
}
func (mcs *MedicineCategoryService) CreateMedicineCategory(newData medicinecategories.MedicineCategory) (*medicinecategories.MedicineCategory, error) {
	result, err := mcs.data.Insert(newData)

	if err != nil {
		return nil, errors.New("Insert Process Failed")
	}

	return result, nil
}
func (mcs *MedicineCategoryService) UpdateMedicineCategory(newData medicinecategories.UpdateMedicineCategory, id int) (bool, error) {
	result, err := mcs.data.Update(newData, id)

	if err != nil {
		return false, errors.New("Update Process Failed")
	}

	return result, nil
}
func (mcs *MedicineCategoryService) DeleteMedicineCategory(id int) (bool, error) {
	result, err := mcs.data.Delete(id)

	if err != nil {
		return false, errors.New("Delete Process Failed")
	}

	return result, nil
}
