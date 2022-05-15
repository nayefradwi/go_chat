package chat

type Chat struct {
	// todo: missing chat id
	User1 User
	User2 User
}

type User struct {
	RefId    int
	Username string
	About    string
	Dp       string
}
