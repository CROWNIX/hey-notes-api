package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func EnvInit() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Err(err).Msg("cannot load env file")
		os.Exit(1)
	}

	AppName = os.Getenv("APP_NAME")
	AppPortInt, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		panic(err)
	}
	AppPort = AppPortInt
	AppStatus = os.Getenv("APP_STATUS")

	DbConnection = os.Getenv("DB_CONNECTION")
	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbDatabase = os.Getenv("DB_DATABASE")
	DbUsername = os.Getenv("DB_USERNAME")
	DbPassword = os.Getenv("DB_PASSWORD")

	log.Info().Msg("config initialization successfully")
}

var (
	AppName string
	AppPort int
	AppStatus string

	DbConnection   string
	DbHost   string
	DbPort   string
	DbDatabase   string
	DbUsername   string
	DbPassword    string
)