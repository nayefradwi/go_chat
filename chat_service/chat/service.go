package chat

type ChatService struct {
	ChatRepo IChatRepo
}

func NewChatService(chatRepo IChatRepo) ChatService {
	return ChatService{
		ChatRepo: chatRepo,
	}
}
