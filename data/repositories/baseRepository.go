package repositories

import (
	"eazyWallet/dataBase"
	"eazyWallet/logger"
	"fmt"
	"gorm.io/gorm"
)

type BaseRepository[T, U any] interface {
	Save(model *T) (*T, error)
	FindById(id *U) (*T, error)
	FindAll() ([]*T, error)
	GetAllBy(name string, value any) ([]*T, error)
}

var dataBaseConnection *gorm.DB

type BaseRepositoryImpl[T, U any] struct{}

func (repository BaseRepositoryImpl[T, U]) Save(model *T) (*T, error) {
	dataBaseConnection = dataBase.DBConnection()
	err := dataBaseConnection.Save(model).Error
	if err != nil {
		logger.ErrorLogger(err)
		return nil, err
	}
	return model, nil
}

func (repository BaseRepositoryImpl[T, U]) FindById(id *U) (*T, error) {
	dataBaseConnection = dataBase.DBConnection()
	var t = new(T)
	err := dataBaseConnection.First(t, id).Error
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (repository BaseRepositoryImpl[T, U]) FindAll() ([]*T, error) {
	dataBaseConnection = dataBase.DBConnection()
	var objects []*T
	err := dataBaseConnection.Find(&objects).Error
	if err != nil {
		return nil, err
	}
	return objects, nil
}

func (repository BaseRepositoryImpl[T, U]) GetAllBy(name string, value any) ([]*T, error) {
	dataBaseConnection = dataBase.DBConnection()
	var objects []*T
	queryString := fmt.Sprintf("%s = ?", name)
	err := dataBaseConnection.Where(queryString, value).Find(&objects).Error
	if err != nil {
		return nil, err
	}
	return objects, nil
}
