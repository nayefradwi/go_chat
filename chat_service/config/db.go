package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateMongoClientAndFetchDatabase(appCtx context.Context) *mongo.Database {
	client := createMongoClient(appCtx)
	return client.Database(DatabaseName)
}

func createMongoClient(appCtx context.Context) *mongo.Client {
	client, err := mongo.Connect(appCtx, options.Client().ApplyURI(DbConnection))
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err.Error())
	}
	log.Print("connected to chat_service database successfully")
	return client
}
