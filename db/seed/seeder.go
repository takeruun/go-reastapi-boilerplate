package main

import (
	"app/config"
	"app/db/seed/seeders"
	"fmt"
)

func main() {
	db := config.NewDB()

	fmt.Println("== Start seed ==")

	if err := seeders.UserSeeds(db); err != nil {
		panic(err.Error())
	}

	if err := seeders.TodoSeeds(db); err != nil {
		panic(err.Error())
	}

	fmt.Println("== Complete seed ==")
}
