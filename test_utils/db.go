package test_utils

import (
	"app/config"
	"errors"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB(t *testing.T) *config.DB {
	c := config.NewConfig()
	db, err := gorm.Open(
		mysql.Open(c.DB.Username+":"+c.DB.Password+"@tcp("+c.DB.Host+")/"+c.DB.DBName+"?charset=utf8&parseTime=True&loc=Local"),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		},
	)
	if err != nil {
		t.Fatal(err)
		panic("fail setup db")
	}

	return &config.DB{DB: db}
}

func RunInTransaction(db *gorm.DB, fn func(db *gorm.DB)) error {
	if err := db.Transaction(func(tx *gorm.DB) error {
		defer func() {
			tx.Rollback()
		}()
		fn(tx)
		return nil
	}); err == nil {
		return errors.New("Failed to Rollback")
	}
	return nil
}
