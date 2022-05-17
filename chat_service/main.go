package main

import (
	"context"
	"log"
	"net/http"

	"github.com/nayefradwi/go_chat/chat_service/config"
	"github.com/nayefradwi/go_chat/chat_service/consumer"
)

func main() {
	config.SetUpEnvironment()
	appCtx := context.Background()
	db := config.CreateMongoClientAndFetchDatabase(appCtx)
	defer db.Client().Disconnect(appCtx)
	consumerGroup := consumer.NewConsumer(config.BrokerList)
	defer consumerGroup.Close()
	r := setupServer(db)
	address := config.Address
	log.Printf("server starting on: %s", address)
	http.ListenAndServe(address, r)
}
