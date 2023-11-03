package data

import "gorm.io/gorm"

type Transaction struct {
	*gorm.Model
	UserID        uint   `gorm:"user_id"`
	TransactionID string `gorm:"column:transaction_id;unique;type:varchar(150)"`
	SnapUrl       string `gorm:"column:snap_url;type:varchar(255)"`
	Status        string `gorm:"column:status;type:varchar(100)"`
	Amount        int    `gorm:"column:amount;type:bigint"`
}
