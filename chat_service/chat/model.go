package chat

type Chat struct {
	// todo: missing chat id
	User1 User `json:"user1" bson:"user1"`
	User2 User
}

type User struct {
	UserRefId int
	Username  string
	About     string
	Dp        string
}
