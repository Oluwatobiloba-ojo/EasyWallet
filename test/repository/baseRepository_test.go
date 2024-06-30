package repository

import (
	"eazyWallet/data/models"
	"eazyWallet/data/repositories"
	"eazyWallet/util/config"
	"log"
	"testing"
)

func TestThatBaseRepositoryCanSave(t *testing.T) {
	config.Load("../../.env")
	repository := repositories.BaseRepositoryImpl[models.Account, uint64]{}
	userAccount := models.Account{AccountBalance: 10000, AccountNumber: "12345"}
	account, err := repository.Save(&userAccount)
	if err != nil {
		t.Error("Error occurred due to that could not create account ", err)
	}
	log.Println(account)
}

func TestThatBaseRepositoryCanFindById(t *testing.T) {
	config.Load("../../.env")
	repository := repositories.BaseRepositoryImpl[models.Account, uint64]{}
	userAccount := models.Account{AccountBalance: 10000, AccountNumber: "12345"}
	account, err := repository.Save(&userAccount)
	if err != nil {
		t.Error("Error occurred due to that could not create account ", err)
	}
	newAccount, err := repository.FindById(&userAccount.ID)
	if err != nil {
		t.Error("Error occurred due to find by id ", err)
	}
	if account.ID != newAccount.ID {
		t.Errorf("Wrong")
	}
}

func TestThatBaseRepositoryFindAll(t *testing.T) {
	config.Load("../../.env")
	repository := repositories.BaseRepositoryImpl[models.Account, uint64]{}
	objects, err := repository.FindAll()
	if err != nil {
		t.Error("Error occurred dur to find all ", err)
	}
	log.Println(&objects)
}

func TestGetAllByInTheBaseRepository(t *testing.T) {
	config.Load("../../.env")
	repository := repositories.BaseRepositoryImpl[models.Transaction, uint64]{}
	_, err := repository.Save(&models.Transaction{AccountId: 1})
	if err != nil {
		t.Error("Error saving a transaction ", err)
	}
	_, err = repository.Save(&models.Transaction{AccountId: 1})
	if err != nil {
		t.Error("Error saving transaction ", err)
	}
	_, err = repository.Save(&models.Transaction{AccountId: 1})
	if err != nil {
		t.Error("Error saving transaction ", err)
	}
	transactions, err := repository.GetAllBy("account_id", 1)
	if err != nil {
		t.Error("Error getting all transactions ", err)
	}
	log.Println(transactions)
}
