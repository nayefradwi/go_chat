package main

import (
	"gochat-user-service/auth"
	"gochat-user-service/friendRequest"
	"gochat-user-service/user"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v4/pgxpool"
)

func SetupServer(dbPool *pgxpool.Pool) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.SetHeader("Content-Type", "application/json"))
	setupUserRoute(r, dbPool)
	setupFriendRequestsRoute(r, dbPool)
	return r
}

func setupUserRoute(r *chi.Mux, dbPool *pgxpool.Pool) {
	userRouter := chi.NewMux()
	userService := user.NewUserService(user.UserRepo{
		Db: dbPool,
	})
	userRouter.Post("/login", userService.Login)
	userRouter.Post("/register", userService.Register)
	userRouter.With(auth.AuthorizeHeaderMiddleware).Get("/user", userService.GetUserById)
	r.Mount("/users", userRouter)
}

func setupFriendRequestsRoute(r *chi.Mux, dbPool *pgxpool.Pool) {
	friendRequestRouter := chi.NewMux()
	friendRequestRouter.Use(auth.AuthorizeHeaderMiddleware)
	friendRequestService := friendrequest.NewFriendRequestService(
		friendrequest.FriendRequestRepo{
			Db: dbPool,
		},
	)
	friendRequestRouter.Post("/{id}/acceptance", friendRequestService.AcceptRequest)
	friendRequestRouter.Get("/", friendRequestService.GetFriendRequests)
	friendRequestRouter.Post("/", friendRequestService.SendFriendRequest)
	friendRequestRouter.Post("/{id}/rejection", friendRequestService.RejectRequest)
	r.Mount("/friend-requests", friendRequestRouter)
}
