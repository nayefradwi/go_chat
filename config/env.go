package config

import (
	"github.com/joho/godotenv"
	"log"
)

const (
	ADDRESS_KEY = "HOST_ADDRESS"
	PORT_KEY    = "PORT"
)

func SetUpEnvironment() {
	if err := godotenv.Load(".local.env"); err != nil {
		log.Fatalf("failed to load environment %s", err)
	}
}
