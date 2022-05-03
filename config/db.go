package config

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

func SetUpDatabaseConnection() *pgxpool.Pool {
	dbPool, err := pgxpool.Connect(context.Background(), DbConnection)
	if err != nil {
		log.Fatalf("failed to set up db connection: %s", err)
	}
	log.Print("connected to database successfully")
	return dbPool
}
