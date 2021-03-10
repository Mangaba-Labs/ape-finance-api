package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func SetupEnvVars() {
	if os.Getenv("ENV") != "PRODUCTION" {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatal("Error loading .env file!")
		}
	}
}
