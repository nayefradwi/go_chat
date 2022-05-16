package friendrequest

import (
	"encoding/json"
	"github.com/nayefradwi/go_chat/user_service/middleware"
	"github.com/nayefradwi/go_chat_common"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type FriendRequestService struct {
	repo IFriendRequestRepo
}

func NewFriendRequestService(friendRequestRepo FriendRequestRepo) FriendRequestService {
	return FriendRequestService{
		repo: friendRequestRepo,
	}
}

func (service FriendRequestService) AcceptRequest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value(middleware.UserIdKey{}).(int)
	requestIdString := chi.URLParam(r, "id")
	requestId, _ := strconv.Atoi(requestIdString)
	err := service.repo.AcceptRequest(ctx, userId, requestId)
	if err != nil {
		gochatcommon.WriteErrorResponse(w, err)
		return
	}
	gochatcommon.WriteEmptySuccessResponse(w, "friend request accepted")
}

func (service FriendRequestService) GetFriendRequests(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value(middleware.UserIdKey{}).(int)
	friendRequests, err := service.repo.GetFriendRequests(ctx, userId)
	if err != nil {
		gochatcommon.WriteErrorResponse(w, err)
		return
	}
	json.NewEncoder(w).Encode(friendRequests)
}

func (service FriendRequestService) SendFriendRequest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userRequestingId := ctx.Value(middleware.UserIdKey{}).(int)
	userRequestedIdString := r.URL.Query().Get("requested-user")
	userRequestedId, _ := strconv.Atoi(userRequestedIdString)
	err := service.repo.SendFriendRequest(ctx, userRequestingId, userRequestedId)
	if err != nil {
		gochatcommon.WriteErrorResponse(w, err)
		return
	}
	gochatcommon.WriteEmptyCreatedResponse(w, "friend request sent")
}

func (service FriendRequestService) RejectRequest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value(middleware.UserIdKey{}).(int)
	requestIdString := chi.URLParam(r, "id")
	requestId, _ := strconv.Atoi(requestIdString)
	err := service.repo.RejectFriendRequest(ctx, userId, requestId)
	if err != nil {
		gochatcommon.WriteErrorResponse(w, err)
		return
	}
	gochatcommon.WriteEmptySuccessResponse(w, "friend request rejected")
}

func (service FriendRequestService) GetSentFriendRequests(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value(middleware.UserIdKey{}).(int)
	friendRequests, err := service.repo.GetSentFriendRequests(ctx, userId)
	if err != nil {
		gochatcommon.WriteErrorResponse(w, err)
		return
	}
	json.NewEncoder(w).Encode(friendRequests)
}
