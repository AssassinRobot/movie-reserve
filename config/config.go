package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

var CurrentENV string 

func configsDirPath() string {
	_, f, _, ok := runtime.Caller(0)
	if !ok {
		panic("Error in generating env dir")
	}

	return filepath.Dir(f)
}

func GetConfigs() (string, string) {
	err := godotenv.Load(configsDirPath() + "/../" + ".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	CurrentENV = os.Getenv("ENV")
	serverPort := os.Getenv("SERVER_PORT")
	DBDsn := os.Getenv("DB_DSN")

	return serverPort, DBDsn
}
