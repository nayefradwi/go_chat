package main

import (
	"context"
	"log"
	"net/http"

	"github.com/nayefradwi/go_chat/user_service/config"
)

func main() {
	config.SetUpEnvironment()
	appCtx := context.Background()
	dbPool := config.SetUpDatabaseConnection(appCtx)
	defer dbPool.Close()
	r, producers := SetupServer(dbPool)
	defer producers.close()
	address := config.Address
	log.Printf("server starting on: %s", address)
	http.ListenAndServe(address, r)
}
