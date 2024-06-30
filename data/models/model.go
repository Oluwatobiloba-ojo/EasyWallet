package models

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	ID             uint64  `gorm:"primaryKey"`
	AccountNumber  string  `gorm:"account_number"`
	AccountBalance float64 `gorm:"account_balance"`
	UserId         uint64  `gorm:"user_id"`
}

type User struct {
	gorm.Model
	ID          uint64 `gorm:"primaryKey"`
	FirstName   string `gorm:"first_name"`
	LastName    string `gorm:"last_name"`
	DateOfBirth string `gorm:"date_of_birth"`
}

type Transaction struct {
	gorm.Model
	ID            uint64  `gorm:"primaryKey"`
	Amount        float64 `gorm:"amount"`
	AccountId     uint64  `gorm:"account_id"`
	Description   string  `gorm:"description"`
	RecipientName string  `gorm:"recipient"`
}
