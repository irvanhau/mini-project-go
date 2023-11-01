package service

import (
	medicalcheckups "MiniProject/features/medical_checkups"
	"errors"
)

type MedicalCheckupService struct {
	data medicalcheckups.MedicalCheckupDataInterface
}

func New(data medicalcheckups.MedicalCheckupDataInterface) medicalcheckups.MedicalCheckupServiceInterface {
	return &MedicalCheckupService{
		data: data,
	}
}

func (mcs *MedicalCheckupService) GetMedicalCheckups() ([]medicalcheckups.MedicalCheckupInfo, error) {
	result, err := mcs.data.GetAll()

	if err != nil {
		return nil, errors.New("Get All Process Failed")
	}

	return result, nil
}
func (mcs *MedicalCheckupService) GetMedicalCheckup(id int) ([]medicalcheckups.MedicalCheckupInfo, error) {
	result, err := mcs.data.GetByID(id)

	if err != nil {
		return nil, errors.New("Get By ID Process Failed")
	}

	return result, nil
}
func (mcs *MedicalCheckupService) CreateMedicalCheckup(newData medicalcheckups.MedicalCheckup) (*medicalcheckups.MedicalCheckup, error) {
	result, err := mcs.data.Insert(newData)

	if err != nil {
		return nil, errors.New("Insert Process Failed")
	}

	return result, nil
}
func (mcs *MedicalCheckupService) UpdateMedicalCheckup(newData medicalcheckups.UpdateMedicalCheckup, id int) (bool, error) {
	result, err := mcs.data.Update(newData, id)

	if err != nil {
		return false, errors.New("Update Process Failed")
	}

	return result, nil
}
func (mcs *MedicalCheckupService) DeleteMedicalCheckup(id int) (bool, error) {
	result, err := mcs.data.Delete(id)

	if err != nil {
		return false, errors.New("Delete Process Failed")
	}

	return result, nil
}
