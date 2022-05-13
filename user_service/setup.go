package main

import (
	"github.com/nayefradwi/go_chat/user_service/auth"
	"github.com/nayefradwi/go_chat/user_service/config"
	"github.com/nayefradwi/go_chat/user_service/friendRequest"
	"github.com/nayefradwi/go_chat/user_service/producer"
	"github.com/nayefradwi/go_chat/user_service/user"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ServerProducers struct {
	userProducer producer.IUserProducer
}

func SetupServer(dbPool *pgxpool.Pool) (*chi.Mux, *ServerProducers) {
	producers := &ServerProducers{}
	r := chi.NewRouter()
	r.Use(middleware.SetHeader("Content-Type", "application/json"))
	setupUserRoute(r, dbPool, producers)
	setupFriendRequestsRoute(r, dbPool)
	return r, producers
}

func setupUserRoute(r *chi.Mux, dbPool *pgxpool.Pool, producers *ServerProducers) {
	userRouter := chi.NewMux()
	userProducer := producer.NewUserProducer([]string{config.Broker0})
	producers.userProducer = userProducer
	userService := user.NewUserService(user.UserRepo{
		Db:           dbPool,
		UserProducer: userProducer,
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
	friendRequestRouter.Get("/sent-requests", friendRequestService.GetSentFriendRequests)
	r.Mount("/friend-requests", friendRequestRouter)
}

func (producers *ServerProducers) close() {
	if userProducer, ok := producers.userProducer.(*producer.UserProducer); ok {
		userProducer.Close()
	}
}
