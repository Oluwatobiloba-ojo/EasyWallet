package models

import (
	"eazyWallet/util/constant"
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
	Status        string    `gorm:"status"`
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	t.Status = constant.PENDING
	return
}
