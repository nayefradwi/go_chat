package chat

import "go.mongodb.org/mongo-driver/bson/primitive"

// todo: add timestamp to sort by last message received on this
type Chat struct {
	Id    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	User1 User               `json:"user1" bson:"user1"`
	User2 User               `json:"user2" bson:"user2"`
}

type User struct {
	UserRefId int    `json:"userRefId" bson:"userRefId"`
	Username  string `json:"username" bson:"username"`
	About     string `json:"about,omitempty" bson:"about,omitempty"`
	Dp        string `json:"dp,omitempty" bson:"dp,omitempty"`
}
