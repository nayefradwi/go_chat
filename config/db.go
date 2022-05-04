package config

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

func SetUpDatabaseConnection(ctx context.Context) *pgxpool.Pool {
	dbPool, err := pgxpool.Connect(ctx, DbConnection)
	if err != nil {
		log.Fatalf("failed to set up db connection: %s", err)
	}
	log.Print("connected to database successfully")
	return dbPool
}
