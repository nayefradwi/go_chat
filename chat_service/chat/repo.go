package chat

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IChatRepo interface {
	getChats(ctx context.Context, userRefId int) []Chat
}

type ChatRepo struct {
	ChatCollection *mongo.Collection
}

func (repo ChatRepo) getChats(ctx context.Context, userRefId int) []Chat {
	var chats []Chat
	cursor, err := repo.ChatCollection.Find(ctx, bson.D{})
	if err != nil {
		return make([]Chat, 0)
	}
	cursor.All(ctx, &chats)
	return chats
}
