package repositories

import (
	"eazyWallet/data/models"
)

type UserRepository interface {
	BaseRepository[models.User, uint]
	GetByPhoneNumber(number string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
}

type UserRepositoryImpl struct {
	*BaseRepositoryImpl[models.User, uint]
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{
		&BaseRepositoryImpl[models.User, uint]{},
	}
}

func (service *UserRepositoryImpl) GetByPhoneNumber(number string) (*models.User, error) {
	user, err := service.GetBy("phone_number", number)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (service *UserRepositoryImpl) GetByEmail(email string) (*models.User, error) {
	user, err := service.GetBy("email", email)
	if err != nil {
		return nil, err
	}
	return user, err
}
