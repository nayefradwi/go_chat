package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nayefradwi/go_chat/chat_service/chat"
	chatServiceMiddleware "github.com/nayefradwi/go_chat/chat_service/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

func setupServer(db *mongo.Database) *chi.Mux {
	r := chi.NewMux()
	r.Use(middleware.SetHeader("Content-Type", "application/json"))
	r.Use(chatServiceMiddleware.AuthorizeHeaderMiddleware)
	setUpChatRoutes(r, db)
	return r
}

func setUpChatRoutes(r *chi.Mux, db *mongo.Database) {
	chatRouter := chi.NewMux()
	chatService := chat.NewChatService(chat.ChatRepo{
		ChatCollection: db.Collection("chats"),
	})
	chatRouter.Get("/", chatService.GetChats)
	r.Mount("/chats", chatRouter)
}
