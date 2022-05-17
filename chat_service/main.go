package main

import (
	"context"
	"log"
	"net/http"

	"github.com/nayefradwi/go_chat/chat_service/config"
)

func main() {
	config.SetUpEnvironment()
	appCtx := context.Background()
	db := config.CreateMongoClientAndFetchDatabase(appCtx)
	defer db.Client().Disconnect(appCtx)
	r := setupServer(db)
	address := config.Address
	log.Printf("server starting on: %s", address)
	http.ListenAndServe(address, r)
}
