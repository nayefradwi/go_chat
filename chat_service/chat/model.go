package chat

import "go.mongodb.org/mongo-driver/bson/primitive"

type Chat struct {
	Id    primitive.ObjectID `json:"id" bson:"_id"`
	User1 User               `json:"user1" bson:"user1"`
	User2 User               `json:"user2" bson:"user2"`
}

type User struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	UserRefId int                `json:"userRefId" bson:"userRefId"`
	Username  string             `json:"username" bson:"username"`
	About     string             `json:"about,omitempty" bson:"about,omitempty"`
	Dp        string             `json:"dp,omitempty" bson:"dp,omitempty"`
}
