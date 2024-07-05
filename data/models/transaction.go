package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID            uuid.UUID `gorm:"primaryKey"`
	Amount        float64   `gorm:"amount"`
	AccountId     uint64    `gorm:"account_id"`
	Description   string    `gorm:"description"`
	RecipientName string    `gorm:"recipient"`
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	return
}
