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
	Host       string
	Username   string
	Password   string
	DBName     string
	Connection *gorm.DB
}

func NewDB() *DB {
	c := NewConfig()
	return newDB(&DB{
		Host:     c.DB.Host,
		Username: c.DB.Username,
		Password: c.DB.Password,
		DBName:   c.DB.DBName,
	})
}

func newDB(d *DB) *DB {
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
		mysql.Open(d.Username+":"+d.Password+"@tcp("+d.Host+")/"+d.DBName+"?charset=utf8&parseTime=True&loc=Local"),
		&gorm.Config{Logger: newLogger},
	)
	db.Logger = db.Logger.LogMode(logger.Info)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("DB接続成功")
	d.Connection = db
	return d
}

// start transaction
func (db *DB) Begin() *gorm.DB {
	return db.Connection.Begin()
}

func (db *DB) Connect() *gorm.DB {
	return db.Connection
}
