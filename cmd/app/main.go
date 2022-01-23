package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mospyx/data-accounting/api"
	"github.com/mospyx/data-accounting/pkg/database"
	"github.com/mospyx/data-accounting/pkg/migrate"
)

//todo: db backup manually/auto
//todo: logout
//todo: generate password handler
//todo: refresh token
//todo: one-time token for create admin

func main() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println(err)
		return
	}

	if err := database.InitDB(); err != nil {
		fmt.Println(err)
		return
	}

	if err := migrate.AutoMigrate(); err != nil {
		fmt.Println(err)
		return
	}

	if err := api.Start(); err != nil {
		fmt.Println(err)
		return
	}
}
