package chat

import (
	"encoding/json"
	"net/http"

	"github.com/nayefradwi/go_chat/chat_service/middleware"
)

type ChatService struct {
	ChatRepo IChatRepo
}

func NewChatService(chatRepo IChatRepo) ChatService {
	return ChatService{
		ChatRepo: chatRepo,
	}
}

func (service ChatService) GetChats(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value(middleware.UserIdKey{}).(int)
	chats := service.ChatRepo.getChats(ctx, userId)
	json.NewEncoder(w).Encode(chats)
}
