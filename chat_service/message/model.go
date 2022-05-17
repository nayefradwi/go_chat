package message

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	TEXT  = iota
	IMAGE = iota
	VIDEO = iota
	FILE  = iota
	URL   = iota
)

type Message struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	ChatId      primitive.ObjectID `json:"chatId" bson:"chatId"`
	SentBy      int                `json:"sentBy" bson:"sentBy"`
	Content     string             `json:"content" bson:"content"`
	ContentType string             `json:"contentType" bson:"contentType"`
	SentAt      time.Time          `json:"sentAt" bson:"sentAt"`
	IsRead      bool               `json:"isRead" bson:"isRead"`
	ReadAt      time.Time          `json:"readAt" bson:"readAt"`
}
