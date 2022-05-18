package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/nayefradwi/go_chat/chat_service/config"
	"github.com/nayefradwi/go_chat/chat_service/consumer"
)

func main() {
	appCtx := generateAppContext()
	db := config.CreateMongoClientAndFetchDatabase(appCtx)
	defer db.Client().Disconnect(appCtx)
	consumerClient := consumer.NewConsumerClient(config.BrokerList)
	defer consumerClient.Close()
	appCtx = context.WithValue(appCtx, consumer.ConsumerClientKey{}, consumerClient)
	r := setupServer(appCtx, db)
	address := config.Address
	log.Printf("server starting on: %s", address)
	http.ListenAndServe(address, r)
}

func generateAppContext() context.Context {
	sigtermChan := make(chan os.Signal, 1)
	signal.Notify(sigtermChan, syscall.SIGINT, syscall.SIGTERM)
	config.SetUpEnvironment()
	appCtx, cancel := context.WithCancel(context.Background())
	go cleanUpRoutinesWithTermination(sigtermChan, cancel)
	return appCtx
}

func cleanUpRoutinesWithTermination(sigtermChan chan os.Signal, cancel context.CancelFunc) {
	log.Print("listening to SIGTERM and SIGINIT")
	<-sigtermChan
	log.Print("server terminated cleaning up go routines")
	cancel()
	os.Exit(1)
}
