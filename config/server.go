package config

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v4/pgxpool"
)

func SetupServer(dbPool *pgxpool.Pool) *chi.Mux {
	r := chi.NewRouter()
	return r
}
