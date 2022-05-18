package main

import (
	"context"

	"github.com/Shopify/sarama"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nayefradwi/go_chat/chat_service/chat"
	"github.com/nayefradwi/go_chat/chat_service/consumer"
	chatServiceMiddleware "github.com/nayefradwi/go_chat/chat_service/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

func setupServer(ctx context.Context, db *mongo.Database) *chi.Mux {
	r := chi.NewMux()
	r.Use(middleware.SetHeader("Content-Type", "application/json"))
	r.Use(chatServiceMiddleware.AuthorizeHeaderMiddleware)
	setUpChatRoutes(ctx, r, db)
	return r
}

func setUpChatRoutes(ctx context.Context, r *chi.Mux, db *mongo.Database) {
	chatsCollection := db.Collection("chats")
	consumerClient := ctx.Value(consumer.ConsumerClientKey{}).(sarama.ConsumerGroup)
	consumer := consumer.NewConsumer(ctx, []string{"NewFriendship"}, consumerClient)
	chatRouter := chi.NewMux()
	chatService := chat.NewChatService(chat.NewChatRepo(ctx, chatsCollection, consumer))
	chatRouter.Get("/", chatService.GetChats)
	r.Mount("/chats", chatRouter)
}
