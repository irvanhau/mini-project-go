package data

import (
	"MiniProject/features/users"
	"MiniProject/helper"
	"errors"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserData struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.UserDataInterface {
	return &UserData{
		db: db,
	}
}

func (ud *UserData) Register(newData users.User) (*users.User, error) {
	var dbData = new(User)
	dbData.Email = newData.Email
	dbData.IdentityNumber = newData.IdentityNumber
	dbData.FullName = newData.FullName
	dbData.BOD = newData.BOD
	dbData.Address = newData.Address
	dbData.Role = newData.Role

	hashPassword, err := helper.HashPassword(newData.Password)
	if err != nil {
		logrus.Info("Hash Password Error, ", err.Error())
	}
	dbData.Password = hashPassword

	if err := ud.db.Create(dbData).Error; err != nil {
		return nil, err
	}

	return &newData, nil
}

func (ud *UserData) Login(email, password string) (*users.User, error) {
	var dbData = new(User)
	dbData.Email = email

	if err := ud.db.Where("email = ?", dbData.Email).First(dbData).Error; err != nil {
		logrus.Info("DB Error : ", err.Error())
		return nil, err
	}

	passwordBytes := []byte(password)
	err := bcrypt.CompareHashAndPassword([]byte(dbData.Password), passwordBytes)

	if err != nil {
		logrus.Info("Incorrect Password")
		return nil, errors.New("Incorrect Password")
	}

	var result = new(users.User)
	result.ID = dbData.ID
	result.Email = dbData.Email
	result.FullName = dbData.FullName
	result.IdentityNumber = dbData.IdentityNumber
	result.BOD = dbData.BOD
	result.Address = dbData.Address
	result.Role = dbData.Role

	return result, nil
}