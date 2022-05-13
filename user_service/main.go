package main

import (
	"context"
	"log"
	"net/http"

	"github.com/nayefradwi/go_chat/user_service/config"
	"github.com/nayefradwi/go_chat/user_service/producer"
)

func main() {
	config.SetUpEnvironment()
	appCtx := context.Background()
	dbPool := config.SetUpDatabaseConnection(appCtx)
	defer dbPool.Close()
	producer := producer.NewProducer([]string{config.Broker0})
	defer producer.Close()
	r := SetupServer(dbPool, &producer)
	address := config.Address
	log.Printf("server starting on: %s", address)
	http.ListenAndServe(address, r)
}
