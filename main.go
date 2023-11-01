package main

import (
	"MiniProject/configs"
	"MiniProject/routes"
	"fmt"

	dataUser "MiniProject/features/users/data"
	handlerUser "MiniProject/features/users/handler"
	serviceUser "MiniProject/features/users/service"

	dataMedicineCategories "MiniProject/features/medicine_categories/data"
	handlerMedicineCategories "MiniProject/features/medicine_categories/handler"
	serviceMedicineCategories "MiniProject/features/medicine_categories/service"

	dataMedicines "MiniProject/features/medicines/data"
	handlerMedicines "MiniProject/features/medicines/handler"
	serviceMedicines "MiniProject/features/medicines/service"

	dataAppointments "MiniProject/features/appointments/data"
	handlerAppointments "MiniProject/features/appointments/handler"
	serviceAppointments "MiniProject/features/appointments/service"

	dataMedicalCheckups "MiniProject/features/medical_checkups/data"
	handlerMedicalCheckups "MiniProject/features/medical_checkups/handler"
	serviceMedicalCheckups "MiniProject/features/medical_checkups/service"

	dataMedicalCheckupDetails "MiniProject/features/medical_checkup_details/data"
	handlerMedicalCheckupDetails "MiniProject/features/medical_checkup_details/handler"
	serviceMedicalCheckupDetails "MiniProject/features/medical_checkup_details/service"

	"MiniProject/helper"
	"MiniProject/utils/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	config := configs.InitConfig()

	var db = database.InitDB(*config)
	database.Migrate(db)

	userModel := dataUser.New(db)
	jwtInterface := helper.New(config.Secret, config.RefSecret)
	userServices := serviceUser.New(userModel, jwtInterface)
	userHandler := handlerUser.NewHandler(userServices)

	medicineCategoryModel := dataMedicineCategories.New(db)
	medicineCategoryServices := serviceMedicineCategories.New(medicineCategoryModel)
	medicineCategoryHandler := handlerMedicineCategories.NewHandler(medicineCategoryServices)

	medicineModel := dataMedicines.New(db)
	medicineServices := serviceMedicines.New(medicineModel)
	medicineHandler := handlerMedicines.NewHandler(medicineServices)

	appointmentModel := dataAppointments.New(db)
	appointmentServices := serviceAppointments.New(appointmentModel)
	appointmentHandler := handlerAppointments.NewHandler(appointmentServices)

	medicalCheckupModel := dataMedicalCheckups.New(db)
	medicalCheckupServices := serviceMedicalCheckups.New(medicalCheckupModel)
	medicalCheckupHandler := handlerMedicalCheckups.NewHandler(medicalCheckupServices)

	medicalCheckupDetailModel := dataMedicalCheckupDetails.New(db)
	medicalCheckupDetailServices := serviceMedicalCheckupDetails.New(medicalCheckupDetailModel)
	medicalCheckupDetailHandler := handlerMedicalCheckupDetails.NewHandler(medicalCheckupDetailServices)

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		}))

	routes.RouteUser(e, userHandler)
	routes.RouteMedicineCategory(e, medicineCategoryHandler, *config)
	routes.RouteMedicine(e, medicineHandler, *config)
	routes.RouteAppointment(e, appointmentHandler, *config)
	routes.RouteMedicalCheckup(e, medicalCheckupHandler, *config)
	routes.RouteMedicalCheckupDetail(e, medicalCheckupDetailHandler, *config)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.ServerPort)).Error())
}
