package routes

import (
	"MiniProject/configs"
	"MiniProject/features/appointments"
	medicinecategories "MiniProject/features/medicine_categories"
	"MiniProject/features/medicines"
	"MiniProject/features/users"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Echo, uh users.UserHandlerInterface) {
	e.POST("/register", uh.Register())
	e.POST("/login", uh.Login())
}

func RouteMedicineCategory(e *echo.Echo, mch medicinecategories.MedicineCategoryHandlerInterface, cfg configs.ProgramConfig) {
	e.GET("/medicine/categories", mch.GetMedicineCategories(), echojwt.JWT([]byte(cfg.Secret)))
	e.GET("/medicine/categories/:id", mch.GetMedicineCategory(), echojwt.JWT([]byte(cfg.Secret)))
	e.POST("/medicine/categories", mch.CreateMedicineCategory(), echojwt.JWT([]byte(cfg.Secret)))
	e.PUT("/medicine/categories/:id", mch.UpdateMedicineCategory(), echojwt.JWT([]byte(cfg.Secret)))
	e.DELETE("/medicine/categories/:id", mch.DeleteMedicineCategory(), echojwt.JWT([]byte(cfg.Secret)))
}

func RouteMedicine(e *echo.Echo, mh medicines.MedicineHandlerInterface, cfg configs.ProgramConfig) {
	e.GET("/medicines", mh.GetMedicines(), echojwt.JWT([]byte(cfg.Secret)))
	e.GET("/medicines/:id", mh.GetMedicine(), echojwt.JWT([]byte(cfg.Secret)))
	e.POST("/medicines", mh.CreateMedicine(), echojwt.JWT([]byte(cfg.Secret)))
	e.PUT("/medicines/:id", mh.UpdateMedicine(), echojwt.JWT([]byte(cfg.Secret)))
	e.DELETE("/medicines/:id", mh.DeleteMedicine(), echojwt.JWT([]byte(cfg.Secret)))

	e.PUT("/medicines/:id/files", mh.UpdateFileMedicine(), echojwt.JWT([]byte(cfg.Secret)))
	e.PUT("/medicines/:id/photos", mh.UpdatePhotoMedicine(), echojwt.JWT([]byte(cfg.Secret)))
}

func RouteAppointment(e *echo.Echo, ah appointments.AppointmentHandlerInterface, cfg configs.ProgramConfig) {
	e.GET("/appointments", ah.GetAppointments(), echojwt.JWT([]byte(cfg.Secret)))
	e.POST("/appointments", ah.CreateAppointment(), echojwt.JWT([]byte(cfg.Secret)))
	e.GET("/appointments/:id", ah.GetAppointment(), echojwt.JWT([]byte(cfg.Secret)))
	e.PUT("/appointments/:id", ah.UpdateAppointment(), echojwt.JWT([]byte(cfg.Secret)))
	e.DELETE("/appointments/:id", ah.DeleteAppointment(), echojwt.JWT([]byte(cfg.Secret)))
}
