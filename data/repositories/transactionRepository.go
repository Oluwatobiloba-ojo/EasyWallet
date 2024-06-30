package repositories

import "eazyWallet/data/models"

type TransactionRepository interface {
	BaseRepository[models.Transaction, uint64]
	FindAllTransactionByAccountId(id uint64) ([]*models.Transaction, error)
}

type TransactionRepositoryImpl struct {
	*BaseRepositoryImpl[models.Transaction, uint64]
}

func NewTransactionRepositoryImpl() TransactionRepository {
	return TransactionRepositoryImpl{
		&BaseRepositoryImpl[models.Transaction, uint64]{},
	}
}

func (receiver TransactionRepositoryImpl) FindAllTransactionByAccountId(id uint64) ([]*models.Transaction, error) {
	transactions, err := receiver.GetAllBy("account_id = ?", id)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
