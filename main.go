package main

import (
	"log"
	"movie/config"
	"movie/database"
)

func main() {
	_, DBDsn := config.GetConfigs()

	_, err := database.NewPostgresDB(DBDsn)
	if err != nil {
		log.Fatal(err)
	}
}
