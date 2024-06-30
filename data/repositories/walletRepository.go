package repositories

import "eazyWallet/data/models"

type WalletRepository interface {
	BaseRepository[models.Account, uint]
}

type NewEasyWalletRepositoryImpl struct {
	*BaseRepositoryImpl[models.Account, uint]
}

func NewWalletRepository() WalletRepository {
	return &NewEasyWalletRepositoryImpl{
		&BaseRepositoryImpl[models.Account, uint]{},
	}
}
