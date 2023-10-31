package handler

import (
	"MiniProject/features/medicines"
	"MiniProject/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MedicineHandler struct {
	service medicines.MedicineServiceInterface
}

func NewHandler(service medicines.MedicineServiceInterface) medicines.MedicineHandlerInterface {
	return &MedicineHandler{
		service: service,
	}
}

func (mh *MedicineHandler) GetMedicines() echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := mh.service.GetMedicines()

		if err != nil {
			c.Logger().Fatal("Handler : Get All Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}
func (mh *MedicineHandler) GetMedicine() echo.HandlerFunc {
	return func(c echo.Context) error {
		paramID := c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			c.Logger().Fatal("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		result, err := mh.service.GetMedicine(id)

		if err != nil {
			c.Logger().Fatal("Handler : Get By ID Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}
func (mh *MedicineHandler) CreateMedicine() echo.HandlerFunc {
	return func(c echo.Context) error {
		formHeaderPhoto, err := c.FormFile("photo")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Select a file to upload", nil))
		}

		formHeaderFile, err := c.FormFile("file")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Select a file to upload", nil))
		}

		var category_id = c.FormValue("category_id")
		var name = c.FormValue("name")
		var stock = c.FormValue("stock")
		var stock_minimum = c.FormValue("stock_minimum")
		var price = c.FormValue("price")

		cat_id, _ := strconv.Atoi(category_id)
		stockInt, _ := strconv.Atoi(stock)
		stockMinInt, _ := strconv.Atoi(stock_minimum)
		priceInt, _ := strconv.Atoi(price)

		formPhoto, err := formHeaderPhoto.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		formFile, err := formHeaderFile.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		uploadUrlPhoto, err := mh.service.PhotoUpload(medicines.MedicinePhoto{Photo: formPhoto})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		uploadUrlFile, err := mh.service.FileUpload(medicines.MedicineFile{File: formFile})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		var serviceInput = new(medicines.Medicine)
		serviceInput.CategoryID = uint(cat_id)
		serviceInput.Name = name
		serviceInput.Stock = stockInt
		serviceInput.StockMinimum = stockMinInt
		serviceInput.Price = priceInt
		serviceInput.Photo = uploadUrlPhoto
		serviceInput.File = uploadUrlFile

		result, err := mh.service.CreateMedicine(*serviceInput)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		var response = new(InputResponse)
		response.CategoryID = result.CategoryID
		response.Name = result.Name
		response.Stock = result.Stock
		response.StockMinimum = result.StockMinimum
		response.Price = result.Price
		response.Photo = result.Photo
		response.File = result.File

		return c.JSON(http.StatusCreated, helper.FormatResponse("Success", response))
	}
}
func (mh *MedicineHandler) UpdateMedicine() echo.HandlerFunc {
	return func(c echo.Context) error {
		var category_id = c.FormValue("category_id")
		var name = c.FormValue("name")
		var stock = c.FormValue("stock")
		var stock_minimum = c.FormValue("stock_minimum")
		var price = c.FormValue("price")

		cat_id, _ := strconv.Atoi(category_id)
		stockInt, _ := strconv.Atoi(stock)
		stockMinInt, _ := strconv.Atoi(stock_minimum)
		priceInt, _ := strconv.Atoi(price)

		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var serviceInput = new(medicines.UpdateMedicine)
		serviceInput.CategoryID = uint(cat_id)
		serviceInput.Name = name
		serviceInput.Stock = stockInt
		serviceInput.StockMinimum = stockMinInt
		serviceInput.Price = priceInt

		result, err := mh.service.UpdateMedicine(*serviceInput, id)

		if err != nil {
			c.Logger().Fatal("Handler : Update Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}

func (mh *MedicineHandler) UpdateFileMedicine() echo.HandlerFunc {
	return func(c echo.Context) error {
		formHeaderFile, err := c.FormFile("file")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Select a file to upload", nil))
		}

		formFile, err := formHeaderFile.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		uploadUrlFile, err := mh.service.FileUpload(medicines.MedicineFile{File: formFile})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		paramID := c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			c.Logger().Fatal("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		result, err := mh.service.UpdateFileMedicine(uploadUrlFile, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}
func (mh *MedicineHandler) UpdatePhotoMedicine() echo.HandlerFunc {
	return func(c echo.Context) error {
		formHeaderPhoto, err := c.FormFile("photo")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Select a file to upload", nil))
		}

		formPhoto, err := formHeaderPhoto.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		uploadUrlPhoto, err := mh.service.PhotoUpload(medicines.MedicinePhoto{Photo: formPhoto})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		paramID := c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			c.Logger().Fatal("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		result, err := mh.service.UpdatePhotoMedicine(uploadUrlPhoto, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}

func (mh *MedicineHandler) DeleteMedicine() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)

		if err != nil {
			c.Logger().Fatal("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		result, err := mh.service.DeleteMedicine(id)

		if err != nil {
			c.Logger().Fatal("Handler : Delete Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusNoContent, helper.FormatResponse("Success", result))
	}
}
