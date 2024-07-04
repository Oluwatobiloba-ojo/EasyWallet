package config

import (
	"github.com/lpernett/godotenv"
	"log"
	"os"
)

var (
	DatabaseUsername       string
	DatabasePassword       string
	DatabaseName           string
	DatabaseHost           string
	PaystackSecretKey      string
	PaystackTransactionUrl string
	MonnifyContractCode    string
	MonnifySecretKey       string
	MonnifyApiKey          string
	MonnifyInitUrl         string
)

func Load(path string) {
	err := godotenv.Load(path)
	if err != nil {
		log.Println("Failed to load file ", err)
		return
	}
	mapConfigurations()
}

func mapConfigurations() {
	mapDataBaseConfiguration()
	mapPaystackConfiguration()
	mapMonnifyCofiguration()
}

func mapDataBaseConfiguration() {
	DatabaseHost = os.Getenv("DB_HOST")
	DatabaseName = os.Getenv("DB_NAME")
	DatabasePassword = os.Getenv("DB_PASSWORD")
	DatabaseUsername = os.Getenv("DB_USERNAME")
}

func mapPaystackConfiguration() {
	PaystackSecretKey = os.Getenv("PAYSTACK_SECRET_KEY")
	PaystackTransactionUrl = os.Getenv("PAYSTACK_INITIALIZE_TRANSACTION_URL")
}

func mapMonnifyCofiguration() {
	MonnifyContractCode = os.Getenv("MONIFY_CONTRACT_CODE")
	MonnifySecretKey = os.Getenv("MONNIFY_SECRET_KEY")
	MonnifyApiKey = os.Getenv("MONNIFY_API_KEY")
	MonnifyInitUrl = os.Getenv("MONNIFY_INIT_URL")
}
