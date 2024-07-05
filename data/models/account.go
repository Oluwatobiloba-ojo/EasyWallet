package models

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	ID             uint64  `gorm:"primaryKey"`
	AccountNumber  string  `gorm:"account_number"`
	AccountBalance float64 `gorm:"account_balance"`
	Password       string  `gorm:"password"`
	UserId         uint64  `gorm:"user_id"`
}
