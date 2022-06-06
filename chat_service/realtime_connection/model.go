package realtimeconnection

import "github.com/nayefradwi/go_chat/chat_service/message"

var ChatConnections map[string]ChatConn = make(map[string]ChatConn)
var ConnectedUsers map[string]UserConn = make(map[string]UserConn)

type UserConn struct {
	UserRefId      int
	Conn           string
	MessageChannel chan message.Message
}

type ChatConn struct {
	ChatId string
	User1  UserConn
	User2  UserConn
}
