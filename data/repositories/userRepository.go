package repositories

import "eazyWallet/data/models"

type UserRepository interface {
	BaseRepository[models.User, uint]
}

type UserRepositoryImpl struct {
	*BaseRepositoryImpl[models.User, uint]
}

func NewUserRepository() UserRepository {
	return UserRepositoryImpl{
		&BaseRepositoryImpl[models.User, uint]{},
	}
}
