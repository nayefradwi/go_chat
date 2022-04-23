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
	setupUserRoute(r, dbPool)
	return r
}

func setupUserRoute(r *chi.Mux, dbPool *pgxpool.Pool) {
	userRouter := chi.NewMux()
	userService := user.NewUserService(user.UserRepo{
		Db: dbPool,
	})
	userRouter.Post("/login", userService.Login)
	userRouter.Post("/register", userService.Register)
	userRouter.Get("/{id}", userService.GetUserById)
	r.Mount("/user", userRouter)
}
