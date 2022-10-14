package config

import (
	"os"
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
}

func NewConfig() *Config {
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

	return c
}
