package chat

import (
	"context"
	"encoding/json"
	"log"

	"github.com/nayefradwi/go_chat/chat_service/consumer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IChatRepo interface {
	getChats(ctx context.Context, userRefId int) []Chat
}

type ChatRepo struct {
	ChatCollection *mongo.Collection
	Consumer       *consumer.Consumer
}

func NewChatRepo(ctx context.Context, chatCollection *mongo.Collection, consumer *consumer.Consumer) ChatRepo {
	repo := ChatRepo{
		ChatCollection: chatCollection,
		Consumer:       consumer,
	}
	go handleConsumedEvent(ctx, &repo)
	return repo
}
func (repo ChatRepo) getChats(ctx context.Context, userRefId int) []Chat {
	chats := make([]Chat, 0)

	userOneOrUserTwo := bson.M{"$or": bson.A{
		bson.M{"user1.userRefId": 9},
		bson.M{"user2.userRefId": 9},
	},
	}
	// userOneOrUserTwo := bson.M{}
	cursor, err := repo.ChatCollection.Find(ctx, userOneOrUserTwo)
	if err != nil {
		return chats
	}
	cursor.All(ctx, &chats)
	return chats
}

func handleConsumedEvent(ctx context.Context, repo *ChatRepo) {
	for {
		eventValue := <-repo.Consumer.ConsumedEvents
		if ctx.Err() != nil {
			log.Print("chat repo stopped consuming events")
		}
		var chat Chat
		json.Unmarshal(eventValue, &chat)
		// todo: make it idemponent?
		insertResult, err := repo.ChatCollection.InsertOne(ctx, chat)
		if err != nil {
			log.Printf("failed to write to database: %s", err.Error())
			continue
		}
		chat.Id = insertResult.InsertedID.(primitive.ObjectID)
	}
}
