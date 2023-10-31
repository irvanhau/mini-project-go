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
}

func NewHandler(service medicinecategories.MedicineCategoryServiceInterface) medicinecategories.MedicineCategoryHandlerInterface {
	return &MedicineCategoryHandler{
		service: service,
	}
}

func (mch *MedicineCategoryHandler) GetMedicineCategories() echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := mch.service.GetMedicineCategories()

		if err != nil {
			c.Logger().Fatal("Handler : Get All Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}
func (mch *MedicineCategoryHandler) GetMedicineCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		paramID := c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			c.Logger().Fatal("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		result, err := mch.service.GetMedicineCategory(id)

		if err != nil {
			c.Logger().Fatal("Handler : Get By ID Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}
func (mch *MedicineCategoryHandler) CreateMedicineCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(InputRequest)
		if err := c.Bind(&input); err != nil {
			c.Logger().Fatal("Handler : Bind Input Error")
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var serviceInput = new(medicinecategories.MedicineCategory)
		serviceInput.Name = input.Name
		serviceInput.Description = input.Description

		result, err := mch.service.CreateMedicineCategory(*serviceInput)

		if err != nil {
			c.Logger().Fatal("Handler : Input Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		var response = new(InputResponse)
		response.Name = result.Name
		response.Description = result.Description

		return c.JSON(http.StatusCreated, helper.FormatResponse("Success", response))
	}
}
func (mch *MedicineCategoryHandler) UpdateMedicineCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		paramID := c.Param("id")
		id, err := strconv.Atoi(paramID)

		if err != nil {
			c.Logger().Fatal("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var input = new(InputRequest)
		if err := c.Bind(&input); err != nil {
			c.Logger().Fatal("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var serviceInput = new(medicinecategories.UpdateMedicineCategory)
		serviceInput.Name = input.Name
		serviceInput.Description = input.Description

		result, err := mch.service.UpdateMedicineCategory(*serviceInput, id)

		if err != nil {
			c.Logger().Fatal("Handler : Update Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}
func (mch *MedicineCategoryHandler) DeleteMedicineCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		paramId := c.Param("id")
		id, err := strconv.Atoi(paramId)
		if err != nil {
			c.Logger().Fatal("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		result, err := mch.service.DeleteMedicineCategory(id)

		if err != nil {
			c.Logger().Fatal("Handler : Delete Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusNoContent, helper.FormatResponse("Success", result))
	}
}
