package friendrequest

import (
	"gochat/errorHandling"

	"github.com/jackc/pgx/v4/pgxpool"
)

type IFriendRequestRepo interface {
	AcceptRequest(requestId int) *errorHandling.BaseError
	GetFriendRequests(userId int) ([]FriendRequest, *errorHandling.BaseError)
	SendFriendRequest(userRequestingId int, userRequestedId int) *errorHandling.BaseError
}

type FriendRequestRepo struct {
	Db *pgxpool.Pool
}

func (repo FriendRequestRepo) AcceptRequest(requestId int) *errorHandling.BaseError {
	return nil
}
func (repo FriendRequestRepo) GetFriendRequests(userId int) ([]FriendRequest, *errorHandling.BaseError) {
	return make([]FriendRequest, 0), nil
}
func (repo FriendRequestRepo) SendFriendRequest(userRequestingId int, userRequestedId int) *errorHandling.BaseError {
	return nil
}
