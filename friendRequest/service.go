package friendrequest

import (
	"encoding/json"
	"gochat/auth"
	"gochat/goChatUtil"
	"net/http"
	"strconv"
)

type FriendRequestService struct {
	repo IFriendRequestRepo
}

func NewFriendRequestService(friendRequestRepo FriendRequestRepo) FriendRequestService {
	return FriendRequestService{
		repo: friendRequestRepo,
	}
}

// todo: accept friend requests
func (service FriendRequestService) AcceptRequest(w http.ResponseWriter, r *http.Request) {

}

func (service FriendRequestService) GetFriendRequests(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(auth.UserIdKey{}).(int)
	friendRequests, err := service.repo.GetFriendRequests(userId)
	if err != nil {
		goChatUtil.WriteErrorResponse(w, err)
		return
	}
	json.NewEncoder(w).Encode(friendRequests)
}

func (service FriendRequestService) SendFriendRequest(w http.ResponseWriter, r *http.Request) {
	userRequestingId := r.Context().Value(auth.UserIdKey{}).(int)
	userRequestedIdString := r.URL.Query().Get("requested-user")
	userRequestedId, _ := strconv.Atoi(userRequestedIdString)
	err := service.repo.SendFriendRequest(userRequestingId, userRequestedId)
	if err != nil {
		goChatUtil.WriteErrorResponse(w, err)
		return
	}
	goChatUtil.WriteEmptyCreatedResponse(w)
}

func (service FriendRequestService) RejectRequest(w http.ResponseWriter, r *http.Request) {

}
