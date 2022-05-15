package main

import (
	"context"

	"github.com/nayefradwi/go_chat/chat_service/config"
)

func main() {
	config.SetUpEnvironment()
	appCtx := context.Background()
	db := config.CreateMongoClientAndFetchDatabase(appCtx)
	defer db.Client().Disconnect(appCtx)
	setupServer(db)
}
