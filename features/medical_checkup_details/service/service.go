package service

import (
	medicalcheckupdetails "MiniProject/features/medical_checkup_details"
	"errors"
)

type MedicalCheckupDetailService struct {
	d medicalcheckupdetails.MedicalCheckupDetailDataInterface
}

func New(data medicalcheckupdetails.MedicalCheckupDetailDataInterface) medicalcheckupdetails.MedicalCheckupDetailServiceInterface {
	return &MedicalCheckupDetailService{
		d: data,
	}
}

func (mcds *MedicalCheckupDetailService) GetMedicalCheckupDetails(idMcu int) ([]medicalcheckupdetails.MedicalCheckupDetailInfo, error) {
	result, err := mcds.d.GetAll(idMcu)

	if err != nil {
		return nil, errors.New("Get All Process Failed")
	}

	return result, nil
}
func (mcds *MedicalCheckupDetailService) GetMedicalCheckupDetail(idMcu, idMcuDetail int) ([]medicalcheckupdetails.MedicalCheckupDetailInfo, error) {
	result, err := mcds.d.GetByID(idMcu, idMcuDetail)

	if err != nil {
		return nil, errors.New("Get By ID Process Failed")
	}

	return result, nil
}
func (mcds *MedicalCheckupDetailService) CreateMedicalCheckupDetail(newData medicalcheckupdetails.MedicalCheckupDetail) (*medicalcheckupdetails.MedicalCheckupDetail, error) {
	result, err := mcds.d.Insert(newData)

	if err != nil {
		return nil, errors.New("Insert Process Failed")
	}

	return result, nil
}
func (mcds *MedicalCheckupDetailService) UpdateMedicalCheckupDetail(newData medicalcheckupdetails.UpdateMedicalCheckupDetail, idMcu, idMcuDetail int) (bool, error) {
	result, err := mcds.d.Update(newData, idMcu, idMcuDetail)

	if err != nil {
		return false, errors.New("Update Process Failed")
	}

	return result, nil
}
func (mcds *MedicalCheckupDetailService) DeleteMedicalCheckupDetail(idMcu, idMcuDetail int) (bool, error) {
	result, err := mcds.d.Delete(idMcu, idMcuDetail)

	if err != nil {
		return false, errors.New("Delete Process Failed")
	}

	return result, nil
}
