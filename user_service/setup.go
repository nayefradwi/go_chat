package main

import (
	"github.com/nayefradwi/go_chat/user_service/config"
	"github.com/nayefradwi/go_chat/user_service/friendRequest"
	"github.com/nayefradwi/go_chat/user_service/producer"
	"github.com/nayefradwi/go_chat/user_service/user"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v4/pgxpool"
	userServiceMiddleware "github.com/nayefradwi/go_chat/user_service/middleware"
)

type ServerProducers struct {
	userProducer          producer.IProducer
	friendRequestProducer producer.IProducer
}

func SetupServer(dbPool *pgxpool.Pool) (*chi.Mux, *ServerProducers) {
	producers := &ServerProducers{}
	r := chi.NewRouter()
	r.Use(middleware.SetHeader("Content-Type", "application/json"))
	setupUserRoute(r, dbPool, producers)
	setupFriendRequestsRoute(r, dbPool, producers)
	return r, producers
}

func setupUserRoute(r *chi.Mux, dbPool *pgxpool.Pool, producers *ServerProducers) {
	userRouter := chi.NewMux()
	userProducer := producer.NewProducer(config.BrokerList)
	producers.userProducer = userProducer
	userService := user.NewUserService(user.UserRepo{
		Db:       dbPool,
		Producer: userProducer,
	})
	userRouter.Post("/login", userService.Login)
	userRouter.Post("/register", userService.Register)
	userRouter.With(userServiceMiddleware.AuthorizeHeaderMiddleware).Get("/user", userService.GetUserById)
	r.Mount("/users", userRouter)
}

func setupFriendRequestsRoute(r *chi.Mux, dbPool *pgxpool.Pool, producers *ServerProducers) {
	friendRequestRouter := chi.NewMux()
	friendRequestProducer := producer.NewProducer(config.BrokerList)
	producers.friendRequestProducer = friendRequestProducer
	friendRequestRouter.Use(userServiceMiddleware.AuthorizeHeaderMiddleware)
	friendRequestService := friendrequest.NewFriendRequestService(
		friendrequest.FriendRequestRepo{
			Db:       dbPool,
			Producer: friendRequestProducer,
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
	producers.userProducer.Close()
	producers.friendRequestProducer.Close()
}
