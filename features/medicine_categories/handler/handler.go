package handler

import (
	medicinecategories "MiniProject/features/medicine_categories"
	"MiniProject/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MedicineCategoryHandler struct {
	service medicinecategories.MedicineCategoryServiceInterface
	jwt     helper.JWTInterface
}

func NewHandler(service medicinecategories.MedicineCategoryServiceInterface, jwt helper.JWTInterface) medicinecategories.MedicineCategoryHandlerInterface {
	return &MedicineCategoryHandler{
		service: service,
		jwt:     jwt,
	}
}

func (mch *MedicineCategoryHandler) GetMedicineCategories() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := mch.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}

		result, err := mch.service.GetMedicineCategories()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success get medicine category", result))
	}
}
func (mch *MedicineCategoryHandler) GetMedicineCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := mch.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}

		paramID := c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid user input", nil))
		}

		result, err := mch.service.GetMedicineCategory(id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success get medicine category", result))
	}
}
func (mch *MedicineCategoryHandler) CreateMedicineCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := mch.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}

		var input = new(InputRequest)
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid user input", nil))
		}

		var serviceInput = new(medicinecategories.MedicineCategory)
		serviceInput.Name = input.Name
		serviceInput.Description = input.Description

		result, err := mch.service.CreateMedicineCategory(*serviceInput)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data", nil))
		}

		var response = new(InputResponse)
		response.Name = result.Name
		response.Description = result.Description

		return c.JSON(http.StatusCreated, helper.FormatResponse("Success create medicine category", response))
	}
}
func (mch *MedicineCategoryHandler) UpdateMedicineCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := mch.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}

		paramID := c.Param("id")
		id, err := strconv.Atoi(paramID)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid user input", nil))
		}

		var input = new(InputRequest)
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid user input", nil))
		}

		var serviceInput = new(medicinecategories.UpdateMedicineCategory)
		serviceInput.Name = input.Name
		serviceInput.Description = input.Description

		result, err := mch.service.UpdateMedicineCategory(*serviceInput, id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success update data", result))
	}
}
func (mch *MedicineCategoryHandler) DeleteMedicineCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := mch.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}

		paramId := c.Param("id")
		id, err := strconv.Atoi(paramId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid user input", nil))
		}

		result, err := mch.service.DeleteMedicineCategory(id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data", nil))
		}

		return c.JSON(http.StatusNoContent, helper.FormatResponse("Success delete data", result))
	}
}
