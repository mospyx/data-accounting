package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mospyx/data_accounting/api"
	"github.com/mospyx/data_accounting/pkg/database"
	"github.com/mospyx/data_accounting/pkg/migrate"
)

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
