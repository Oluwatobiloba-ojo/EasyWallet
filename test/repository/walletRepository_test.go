package repository

import (
	"eazyWallet/data/models"
	"eazyWallet/data/repositories"
	"eazyWallet/util/config"
	"log"
	"testing"
)

func TestThatWalletRepositoryCanOnlySaveWallet(t *testing.T) {
	config.Load("../../.env")
	walletRepository := repositories.NewWalletRepository()
	wallet := models.Account{AccountNumber: "123456", AccountBalance: 10000}
	userWallet, err := walletRepository.Save(&wallet)
	if err != nil {
		t.Error("Error saving wallet ", err)
	}
	log.Println(userWallet.ID)
}
