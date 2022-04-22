package config

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
)

func SetUpDatabaseConnection() *pgxpool.Pool {
	dbPool, err := pgxpool.Connect(context.Background(), os.Getenv(DB_CONNECTION))
	if err != nil {
		log.Fatalf("failed to set up db connection: %s", err)
	}
	log.Print("connected to database successfully")
	return dbPool
}
