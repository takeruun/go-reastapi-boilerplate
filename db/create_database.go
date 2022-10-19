package main

import (
	"app/config"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	config := config.NewConfig()

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: false,         // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,          // Disable color
		},
	)

	db, err := gorm.Open(
		mysql.Open(config.DB.Username+":"+config.DB.Password+"@tcp("+config.DB.Host+")/?charset=utf8&parseTime=True&loc=Local"),
		&gorm.Config{Logger: newLogger},
	)

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("DB接続成功")

	fmt.Println("== Start create database ==")

	err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`;", config.DB.DBName)).Error
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("== Complete Create Database %v ==\n", config.DB.DBName)
}
