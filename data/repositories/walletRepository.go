package repositories

import "eazyWallet/data/models"

type WalletRepository interface {
	BaseRepository[models.Account, uint64]
	FindWalletByAccountNumber(number string) (*models.Account, error)
}

func (repository *NewEasyWalletRepositoryImpl) FindWalletByAccountNumber(number string) (*models.Account, error) {
	account, err := repository.GetBy("account_number", number)
	if err != nil {
		return nil, err
	}
	return account, nil
}

type NewEasyWalletRepositoryImpl struct {
	*BaseRepositoryImpl[models.Account, uint64]
}

func NewWalletRepository() WalletRepository {
	return &NewEasyWalletRepositoryImpl{
		&BaseRepositoryImpl[models.Account, uint64]{},
	}
}
