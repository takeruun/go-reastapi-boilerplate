package seeders

import (
	"app/config"
	"app/entity"
	"fmt"
	"math/rand"
	"strconv"
)

func TodoSeeds(db *config.DB) error {
	for i := 1; i <= 10; i++ {
		todo := entity.Todo{
			Title:       "タイトル_" + strconv.Itoa(i),
			Description: "説明_" + strconv.Itoa(i),
			UserId:      uint64(rand.Intn(4) + 1),
		}

		if err := db.Create(&todo).Error; err != nil {
			fmt.Printf("%+v", err)
		}
	}
	return nil
}
