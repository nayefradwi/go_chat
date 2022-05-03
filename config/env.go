package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	ADDRESS       = "HOST_ADDRESS"
	DB_CONNECTION = "DATABASE_CONNECTION"
	SECRET        = "SECRET"
)

var (
	Secret       string
	Address      string
	DbConnection string
)

func SetUpEnvironment() {
	if err := godotenv.Load(".local.env"); err != nil {
		log.Fatalf("failed to load environment %s", err)
	}
	Secret = os.Getenv(SECRET)
	Address = os.Getenv(ADDRESS)
	DbConnection = os.Getenv(DB_CONNECTION)
}
