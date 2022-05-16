package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

func setupServer(db *mongo.Database) *chi.Mux {
	r := chi.NewMux()
	r.Use(middleware.SetHeader("Content-Type", "application/json"))
	return r
}

func setUpChatRoutes(db *mongo.Database) {
	// todo: missing authentication route
}
