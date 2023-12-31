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
	jwt     helper.JWTInterface
}

func NewHandler(service medicines.MedicineServiceInterface, jwt helper.JWTInterface) medicines.MedicineHandlerInterface {
	return &MedicineHandler{
		service: service,
		jwt:     jwt,
	}
}

func (mh *MedicineHandler) GetMedicines() echo.HandlerFunc {
	return func(c echo.Context) error {

		role := mh.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}

		var paramKategori = c.QueryParam("kategori")
		var name = c.QueryParam("name")
		kategori, _ := strconv.Atoi(paramKategori)

		result, err := mh.service.GetMedicines(kategori, name)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success get medicine", result))
	}
}
func (mh *MedicineHandler) GetMedicine() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := mh.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}

		paramID := c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid user input", nil))
		}

		result, err := mh.service.GetMedicine(id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success get medicine", result))
	}
}
func (mh *MedicineHandler) CreateMedicine() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := mh.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}

		formHeaderPhoto, err := c.FormFile("photo")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Select a photo to upload", nil))
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
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot open foto", nil))
		}

		formFile, err := formHeaderFile.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot open file", nil))
		}

		uploadUrlPhoto, err := mh.service.PhotoUpload(medicines.MedicinePhoto{Photo: formPhoto})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot upload photo", nil))
		}

		uploadUrlFile, err := mh.service.FileUpload(medicines.MedicineFile{File: formFile})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot upload file", nil))
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
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data", nil))
		}

		var response = new(InputResponse)
		response.CategoryID = result.CategoryID
		response.Name = result.Name
		response.Stock = result.Stock
		response.StockMinimum = result.StockMinimum
		response.Price = result.Price
		response.Photo = result.Photo
		response.File = result.File

		return c.JSON(http.StatusCreated, helper.FormatResponse("Success created medicine", response))
	}
}
func (mh *MedicineHandler) UpdateMedicine() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := mh.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
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

		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid user input", nil))
		}

		var serviceInput = new(medicines.UpdateMedicine)
		serviceInput.CategoryID = uint(cat_id)
		serviceInput.Name = name
		serviceInput.Stock = stockInt
		serviceInput.StockMinimum = stockMinInt
		serviceInput.Price = priceInt

		result, err := mh.service.UpdateMedicine(*serviceInput, id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success update data", result))
	}
}

func (mh *MedicineHandler) UpdateFileMedicine() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := mh.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}

		formHeaderFile, err := c.FormFile("file")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Select a file to upload", nil))
		}

		formFile, err := formHeaderFile.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot open file", nil))
		}

		uploadUrlFile, err := mh.service.FileUpload(medicines.MedicineFile{File: formFile})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot upload file", nil))
		}

		paramID := c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid user input", nil))
		}

		result, err := mh.service.UpdateFileMedicine(uploadUrlFile, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success update file", result))
	}
}
func (mh *MedicineHandler) UpdatePhotoMedicine() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := mh.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}

		formHeaderPhoto, err := c.FormFile("photo")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Select a file to upload", nil))
		}

		formPhoto, err := formHeaderPhoto.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot open photo", nil))
		}

		uploadUrlPhoto, err := mh.service.PhotoUpload(medicines.MedicinePhoto{Photo: formPhoto})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot upload photo", nil))
		}

		paramID := c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid user input", nil))
		}

		result, err := mh.service.UpdatePhotoMedicine(uploadUrlPhoto, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success update photo", result))
	}
}

func (mh *MedicineHandler) DeleteMedicine() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := mh.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}

		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid user input", nil))
		}

		result, err := mh.service.DeleteMedicine(id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data", nil))
		}

		return c.JSON(http.StatusNoContent, helper.FormatResponse("Success delete data", result))
	}
}
