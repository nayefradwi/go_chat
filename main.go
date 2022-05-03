package main

import (
	"gochat/config"
	"log"
	"net/http"
)

func main() {
	config.SetUpEnvironment()
	dbPool := config.SetUpDatabaseConnection()
	defer dbPool.Close()
	r := SetupServer(dbPool)
	address := config.Address
	log.Printf("server starting on: %s", address)
	http.ListenAndServe(address, r)
}
