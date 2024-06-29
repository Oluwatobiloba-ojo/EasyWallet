package config

import (
	"github.com/lpernett/godotenv"
	"log"
	"os"
)

var (
	DatabaseUsername string
	DatabasePassword string
	DatabaseName     string
	DatabaseHost     string
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
}

func mapDataBaseConfiguration() {
	DatabaseHost = os.Getenv("DB_HOST")
	DatabaseName = os.Getenv("DB_NAME")
	DatabasePassword = os.Getenv("DB_PASSWORD")
	DatabaseUsername = os.Getenv("DB_USERNAME")
}
