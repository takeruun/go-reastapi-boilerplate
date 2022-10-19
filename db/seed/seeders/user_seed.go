package seeders

import (
	"app/config"
	"app/entity"
	"fmt"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func UserSeeds(db *config.DB) error {
	for i := 1; i <= 5; i++ {
		hash, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		user := entity.User{
			Name:         "ユーザー_" + strconv.Itoa(i),
			Email:        "test" + strconv.Itoa(i) + "@example.com",
			HashPassword: string(hash),
		}

		if err := db.Create(&user).Error; err != nil {
			fmt.Printf("%+v", err)
		}
	}
	return nil
}
