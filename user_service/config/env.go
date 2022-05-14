package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const (
	ADDRESS       = "HOST_ADDRESS"
	DB_CONNECTION = "DATABASE_CONNECTION"
	SECRET        = "SECRET"
	BROKER_LIST   = "BROKER_LIST"
)

var (
	Secret       string
	Address      string
	DbConnection string
	BrokerList   []string
)

func SetUpEnvironment() {
	godotenv.Load(".local.env")
	Secret = os.Getenv(SECRET)
	Address = os.Getenv(ADDRESS)
	DbConnection = os.Getenv(DB_CONNECTION)
	BrokerList = strings.Split(os.Getenv(BROKER_LIST), ";")
}
