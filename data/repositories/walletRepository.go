package repositories

import "eazyWallet/data/models"

type WalletRepository interface {
	BaseRepository[models.Account, uint64]
}

type NewEasyWalletRepositoryImpl struct {
	*BaseRepositoryImpl[models.Account, uint64]
}

func NewWalletRepository() WalletRepository {
	return &NewEasyWalletRepositoryImpl{
		&BaseRepositoryImpl[models.Account, uint64]{},
	}
}
