package database

import (
	"MiniProject/configs"
	"fmt"

	"github.com/labstack/gommon/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(c configs.ProgramConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DBUser,
		c.DBPass,
		c.DBHost,
		c.DBPort,
		c.DBName,
	)
	log.Print(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Error("terjadi kesalahan pada database, error:", err.Error())
		return nil, err
	}

	return db, nil
}
