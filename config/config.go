package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	applicationDir = "app/"
)

type Config struct {
	DB struct {
		Host     string
		Username string
		Password string
		DBName   string
	}
	Routing struct {
		Port string
	}
	SESSION_STORE struct {
		SecretHashKey string
	}
	Mail struct {
		Mode string
		Auth struct {
			Host     string
			Email    string
			Password string
		}
		Addr      string
		FromName  string
		FromEmail string
	}
}

func NewConfig() *Config {
	err := godotenv.Load(os.ExpandEnv("/go/src/" + applicationDir + ".env"))
	if err != nil {
		fmt.Println("env file 読み込み出来ませんでした。")
	}
	c := new(Config)

	goMode := os.Getenv("GO_MODE")
	switch goMode {
	case "development":
		c.DB.DBName = os.Getenv("DB_NAME") + "-dev"
	case "test":
		c.DB.DBName = os.Getenv("DB_NAME") + "-test"
	case "prodcution":
		c.DB.DBName = os.Getenv("DB_NAME")
	default:
		c.DB.DBName = os.Getenv("DB_NAME") + "-dev"
	}

	c.DB.Host = os.Getenv("DB_HOST")
	c.DB.Username = os.Getenv("DB_USER")
	c.DB.Password = os.Getenv("DB_PASSWORD")

	c.SESSION_STORE.SecretHashKey = os.Getenv("SECRET_HASH_KEY")

	c.Routing.Port = "3000"

	c.Mail.Mode = "google"
	c.Mail.Auth.Host = "smtp.gmail.com"
	c.Mail.Auth.Email = os.Getenv("MAIL_AUTH_EMAIL")
	c.Mail.Auth.Password = os.Getenv("MAIL_AUTH_PASSWORD")
	c.Mail.Addr = "smtp.gmail.com:587"
	c.Mail.FromName = os.Getenv("MAIL_FROM_NAME")
	c.Mail.FromEmail = os.Getenv("MAIL_FROM_EMAIL")

	return c
}
