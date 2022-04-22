package main

import (
	"gochat/config"
	"log"
	"net/http"
	"os"
)

func main() {
	config.SetUpEnvironment()
	dbPool := config.SetUpDatabaseConnection()
	r := config.SetupServer(dbPool)
	address := os.Getenv(config.ADDRESS)
	log.Printf("server starting on: %s", address)
	http.ListenAndServe(address, r)
}
