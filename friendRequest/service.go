package friendrequest

import "net/http"

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

// todo: get all friend requests
func (service FriendRequestService) GetFriendRequests(w http.ResponseWriter, r *http.Request) {

}

// todo: send friend requests
func (service FriendRequestService) SendFriendRequest(w http.ResponseWriter, r *http.Request) {

}
