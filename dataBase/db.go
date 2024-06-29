package dataBase

import (
	"database/sql"
	"eazyWallet/logger"
	"eazyWallet/util/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func DBConnection() *gorm.DB {
	err := EnsureDataBaseCreated()
	if err != nil {
		logger.ErrorLogger(err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DatabaseUsername, config.DatabasePassword, config.DatabaseHost, config.DatabaseName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.ErrorLogger(err)
	}
	return db
}

func EnsureDataBaseCreated() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/", config.DatabaseUsername, config.DatabasePassword, config.DatabaseHost)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("failed to connect to database")
		return err
	}
	defer db.Close()
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + config.DatabaseName)
	if err != nil {
		log.Println("Failed to create database ")
		return err
	}
	return nil
}
