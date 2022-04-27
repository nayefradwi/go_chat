package config

import (
	"github.com/joho/godotenv"
	"log"
)

const (
	ADDRESS       = "HOST_ADDRESS"
	DB_CONNECTION = "DATABASE_CONNECTION"
	SECRET        = "SECRET"
)

func SetUpEnvironment() {
	if err := godotenv.Load(".local.env"); err != nil {
		log.Fatalf("failed to load environment %s", err)
	}
}
