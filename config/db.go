package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	*gorm.DB
}

func NewDB() *DB {
	c := NewConfig()
	return &DB{
		connectDB(c),
	}
}

func connectDB(config *Config) *gorm.DB {
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
		mysql.Open(config.DB.Username+":"+config.DB.Password+"@tcp("+config.DB.Host+")/"+config.DB.DBName+"?charset=utf8&parseTime=True&loc=Local"),
		&gorm.Config{Logger: newLogger},
	)
	db.Logger = db.Logger.LogMode(logger.Info)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("DB接続成功")

	return db
}

// start transaction
func (db *DB) Begin() *gorm.DB {
	return db.Begin()
}
