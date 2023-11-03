package service

import (
	"MiniProject/features/medicines"
	"MiniProject/utils/cloudinary"
	"errors"
)

type MedicineService struct {
	data medicines.MedicineDataInterface
}

func New(data medicines.MedicineDataInterface) medicines.MedicineServiceInterface {
	return &MedicineService{
		data: data,
	}
}

func (md *MedicineService) GetMedicines(kategori int, name string) ([]medicines.MedicineInfo, error) {
	result, err := md.data.GetAll(kategori, name)

	if err != nil {
		return nil, errors.New("Get All Failed")
	}

	return result, nil
}

func (md *MedicineService) GetMedicine(id int) ([]medicines.MedicineInfo, error) {
	result, err := md.data.GetByID(id)

	if err != nil {
		return nil, errors.New("Get By ID Failed")
	}

	return result, nil
}

func (md *MedicineService) CreateMedicine(newData medicines.Medicine) (*medicines.Medicine, error) {
	result, err := md.data.Insert(newData)

	if err != nil {
		return nil, errors.New("Insert Process Failed")
	}

	return result, nil
}

func (md *MedicineService) UpdateMedicine(newData medicines.UpdateMedicine, id int) (bool, error) {
	result, err := md.data.Update(newData, id)

	if err != nil {
		return false, errors.New("Update Process Failed")
	}

	return result, nil
}

func (md *MedicineService) DeleteMedicine(id int) (bool, error) {
	result, err := md.data.Delete(id)

	if err != nil {
		return false, errors.New("Delete Process Failed")
	}

	return result, nil
}

func (md *MedicineService) FileUpload(file medicines.MedicineFile) (string, error) {
	uploadUrl, err := cloudinary.FileUploadHelper(file.File)

	if err != nil {
		return "", errors.New("Upload File Failed")
	}

	return uploadUrl, nil
}

func (md *MedicineService) PhotoUpload(file medicines.MedicinePhoto) (string, error) {
	uploadUrl, err := cloudinary.ImageUploadHelper(file.Photo)

	if err != nil {
		return "", errors.New("Upload Photo Failed")
	}

	return uploadUrl, nil
}

func (md *MedicineService) UpdateFileMedicine(file string, id int) (bool, error) {
	result, err := md.data.UpdateFileMedicine(file, id)

	if err != nil {
		return false, errors.New("Update File Medicine Failed")
	}

	return result, nil
}

func (md *MedicineService) UpdatePhotoMedicine(photo string, id int) (bool, error) {
	result, err := md.data.UpdatePhotoMedicine(photo, id)

	if err != nil {
		return false, errors.New("Update Photo Medicine Failed")
	}

	return result, nil
}
