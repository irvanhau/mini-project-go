package service

import (
	"MiniProject/features/medicines"
	"MiniProject/features/medicines/mocks"
	utils "MiniProject/utils/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMedicines(t *testing.T) {
	data := mocks.NewMedicineDataInterface(t)
	cl := utils.NewCloudinaryInterface(t)
	service := New(data, cl)
	medicine := []medicines.MedicineInfo{}

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetAll", 1, "p").Return(medicine, nil).Once()

		result, err := service.GetMedicines(1, "p")

		assert.Nil(t, err)
		assert.NotNil(t, result)
		data.AssertExpectations(t)
		cl.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetAll", 1, "p").Return(nil, errors.New("Get All Failed")).Once()

		result, err := service.GetMedicines(1, "p")
		assert.Error(t, err)
		assert.EqualError(t, err, "Get All Failed")
		assert.Nil(t, result)
	})
}

func TestGetMedicine(t *testing.T) {
	data := mocks.NewMedicineDataInterface(t)
	cl := utils.NewCloudinaryInterface(t)
	service := New(data, cl)
	medicine := []medicines.MedicineInfo{}

	t.Run("Success Get", func(t *testing.T) {
		data.On("GetByID", 1).Return(medicine, nil).Once()

		result, err := service.GetMedicine(1)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		data.AssertExpectations(t)
		cl.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("GetByID", 1).Return(nil, errors.New("Get By ID Failed")).Once()

		result, err := service.GetMedicine(1)
		assert.Error(t, err)
		assert.EqualError(t, err, "Get By ID Failed")
		assert.Nil(t, result)
	})
}

func TestCreateMedicine(t *testing.T) {
	data := mocks.NewMedicineDataInterface(t)
	cl := utils.NewCloudinaryInterface(t)
	service := New(data, cl)
	medicine := medicines.Medicine{
		CategoryID:   1,
		Name:         "Panadol",
		StockMinimum: 5,
		Stock:        10,
		Price:        10000,
		Photo:        "http.cloudinary.com",
		File:         "http.cloudinary.com",
	}

	t.Run("Success Insert", func(t *testing.T) {
		data.On("Insert", medicine).Return(&medicine, nil).Once()

		result, err := service.CreateMedicine(medicine)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, medicine.CategoryID, result.CategoryID)
		assert.Equal(t, medicine.Name, result.Name)
		assert.Equal(t, medicine.StockMinimum, result.StockMinimum)
		assert.Equal(t, medicine.Stock, result.Stock)
		assert.Equal(t, medicine.Price, result.Price)
		assert.Equal(t, medicine.Photo, result.Photo)
		assert.Equal(t, medicine.File, result.File)
		data.AssertExpectations(t)
		cl.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("Insert", medicine).Return(nil, errors.New("Insert Process Failed")).Once()

		result, err := service.CreateMedicine(medicine)
		assert.Error(t, err)
		assert.EqualError(t, err, "Insert Process Failed")
		assert.Nil(t, result)
	})
}

func TestUpdateMedicine(t *testing.T) {
	data := mocks.NewMedicineDataInterface(t)
	cl := utils.NewCloudinaryInterface(t)
	service := New(data, cl)
	medicine := medicines.UpdateMedicine{
		CategoryID:   1,
		Name:         "Panadol",
		StockMinimum: 5,
		Stock:        10,
		Price:        10000,
	}

	t.Run("Success Update", func(t *testing.T) {
		data.On("Update", medicine, 1).Return(true, nil).Once()

		result, err := service.UpdateMedicine(medicine, 1)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, true, result)
		data.AssertExpectations(t)
		cl.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("Update", medicine, 1).Return(false, errors.New("Update Process Failed")).Once()

		result, err := service.UpdateMedicine(medicine, 1)
		assert.Error(t, err)
		assert.EqualError(t, err, "Update Process Failed")
		assert.Equal(t, false, result)
	})
}

func TestDeleteMedicine(t *testing.T) {
	data := mocks.NewMedicineDataInterface(t)
	cl := utils.NewCloudinaryInterface(t)
	service := New(data, cl)

	t.Run("Success Delete", func(t *testing.T) {
		data.On("Delete", 1).Return(true, nil).Once()

		result, err := service.DeleteMedicine(1)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, true, result)
		data.AssertExpectations(t)
		cl.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("Delete", 1).Return(false, errors.New("Delete Process Failed")).Once()

		result, err := service.DeleteMedicine(1)
		assert.Error(t, err)
		assert.EqualError(t, err, "Delete Process Failed")
		assert.Equal(t, false, result)
	})
}

func TestUpdateFileMedicine(t *testing.T) {
	data := mocks.NewMedicineDataInterface(t)
	cl := utils.NewCloudinaryInterface(t)
	service := New(data, cl)

	t.Run("Success Update File", func(t *testing.T) {
		data.On("UpdateFileMedicine", "", 1).Return(true, nil).Once()

		result, err := service.UpdateFileMedicine("", 1)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		data.AssertExpectations(t)
		cl.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("UpdateFileMedicine", "", 1).Return(false, errors.New("Update File Medicine Failed")).Once()

		result, err := service.UpdateFileMedicine("", 1)

		assert.Error(t, err)
		assert.EqualError(t, err, "Update File Medicine Failed")
		assert.Equal(t, false, result)
	})
}

func TestUpdatePhotoMedicine(t *testing.T) {
	data := mocks.NewMedicineDataInterface(t)
	cl := utils.NewCloudinaryInterface(t)
	service := New(data, cl)

	t.Run("Success Update Photo", func(t *testing.T) {
		data.On("UpdatePhotoMedicine", "", 1).Return(true, nil).Once()

		result, err := service.UpdatePhotoMedicine("", 1)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		data.AssertExpectations(t)
		cl.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		data.On("UpdatePhotoMedicine", "", 1).Return(false, errors.New("Update Photo Medicine Failed")).Once()

		result, err := service.UpdatePhotoMedicine("", 1)

		assert.Error(t, err)
		assert.EqualError(t, err, "Update Photo Medicine Failed")
		assert.Equal(t, false, result)
	})
}

// func TestFileUpload(t *testing.T) {
// 	// data := mocks.NewMedicineDataInterface(t)
// 	cloud := cloudinary.NewCloudinaryInterface(t)

// 	t.Run("Success Upload File", func(t *testing.T) {
// 		cloud.On("FileUploadHelper", "file.jpg").Return("Success", nil).Once()

// 		result, err := cloud.FileUploadHelper("file.jpg")

// 		assert.Nil(t, err)
// 		assert.NotNil(t, result)
// 		assert.Equal(t, "Success", result)
// 		cloud.AssertExpectations(t)
// 	})

// 	t.Run("Cloudinary Error", func(t *testing.T) {
// 		cloud.On("FileUploadHelper", "file.jpg").Return("", errors.New("Upload File Failed")).Once()

// 		result, err := cloud.FileUploadHelper("file.jpg")

// 		assert.Error(t, err)
// 		assert.EqualError(t, err, "Upload File Failed")
// 		assert.Equal(t, "", result)
// 	})
// }

// func TestPhotoUpload(t *testing.T) {
// 	data := mocks.NewMedicineDataInterface(t)
// 	service := New(data)
// 	cloud := cloudinary.NewCloudinaryInterface(t)

// 	file, err := os.Open("./MiniProjectKMB5.png")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	medicinePhoto := medicines.MedicinePhoto{
// 		Photo: file,
// 	}

// 	t.Run("Success Upload Photo", func(t *testing.T) {
// 		cloud.On("ImageUploadHelper", "image.jpg").Return("Success", nil).Once()

// 		result, err := service.PhotoUpload(medicinePhoto)

// 		assert.Nil(t, err)
// 		assert.NotNil(t, result)
// 		assert.Equal(t, "Success", result)
// 		cloud.AssertExpectations(t)
// 	})

// 	t.Run("Cloudinary Error", func(t *testing.T) {
// 		cloud.On("ImageUploadHelper", "image.jpg").Return("", errors.New("Upload Photo Failed")).Once()

// 		result, err := cloud.ImageUploadHelper("image.jpg")

// 		assert.Error(t, err)
// 		assert.EqualError(t, err, "Upload Photo Failed")
// 		assert.Equal(t, "", result)
// 	})
// }
