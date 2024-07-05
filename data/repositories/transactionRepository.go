package repositories

import (
	"eazyWallet/data/models"
	"github.com/google/uuid"
)

type TransactionRepository interface {
	BaseRepository[models.Transaction, uuid.UUID]
	FindAllTransactionByAccountId(id uint64) ([]*models.Transaction, error)
}

type TransactionRepositoryImpl struct {
	*BaseRepositoryImpl[models.Transaction, uuid.UUID]
}

func NewTransactionRepositoryImpl() TransactionRepository {
	return TransactionRepositoryImpl{
		&BaseRepositoryImpl[models.Transaction, uuid.UUID]{},
	}
}

func (receiver TransactionRepositoryImpl) FindAllTransactionByAccountId(id uint64) ([]*models.Transaction, error) {
	transactions, err := receiver.GetAllBy("account_id", id)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
