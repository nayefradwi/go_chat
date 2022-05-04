package main

import (
	"context"
	"gochat/config"
	"log"
	"net/http"
)

func main() {
	config.SetUpEnvironment()
	appCtx := context.Background()
	dbPool := config.SetUpDatabaseConnection(appCtx)
	defer dbPool.Close()
	r := SetupServer(dbPool)
	address := config.Address
	log.Printf("server starting on: %s", address)
	http.ListenAndServe(address, r)
}
