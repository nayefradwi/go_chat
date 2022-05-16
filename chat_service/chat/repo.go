package chat

import "go.mongodb.org/mongo-driver/mongo"

type IChatRepo interface {
	getChats(userRefId int) []Chat
}

type ChatRepo struct {
	ChatCollection mongo.Collection
}
