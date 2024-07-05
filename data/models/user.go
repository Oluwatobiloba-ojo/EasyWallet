package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID          uint64 `gorm:"primaryKey"`
	FirstName   string `gorm:"first_name"`
	LastName    string `gorm:"last_name"`
	PhoneNumber string `gorm:"phone_number"`
	Email       string `gorm:"email"`
}
