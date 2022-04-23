package config

import (
	"gochat/user"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v4/pgxpool"
)

func SetupServer(dbPool *pgxpool.Pool) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.SetHeader("Content-Type", "application/json"))
	userService := user.NewUserService(user.UserRepo{})
	r.Get("/user", userService.GetUserById)
	return r
}
