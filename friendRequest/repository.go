package friendrequest

import (
	"context"
	"gochat/errorHandling"

	"github.com/jackc/pgx/v4/pgxpool"
)

type IFriendRequestRepo interface {
	AcceptRequest(requestId int) *errorHandling.BaseError
	GetFriendRequests(userId int) ([]FriendRequestDetails, *errorHandling.BaseError)
	SendFriendRequest(userRequestingId int, userRequestedId int) *errorHandling.BaseError
	RejectFriendRequest(requestId int) *errorHandling.BaseError
}

type FriendRequestRepo struct {
	Db *pgxpool.Pool
}

func (repo FriendRequestRepo) AcceptRequest(requestId int) *errorHandling.BaseError {
	return nil
}
func (repo FriendRequestRepo) GetFriendRequests(userId int) ([]FriendRequestDetails, *errorHandling.BaseError) {
	friendRequests := make([]FriendRequestDetails, 0)
	rows, err := repo.Db.Query(context.Background(), GET_FRIEND_REQUESTS, userId)
	if err != nil {
		return friendRequests, errorHandling.NewBadRequest("failed to load friend requests")
	}
	for rows.Next() {
		var id int
		var username, about, dp string
		err := rows.Scan(&id, &username, &about, &dp)
		if err == nil {
			friendRequests = append(friendRequests, FriendRequestDetails{
				Id:       id,
				Username: username,
				About:    about,
				Dp:       dp,
			})
		}
	}
	return friendRequests, nil
}

func (repo FriendRequestRepo) SendFriendRequest(userRequestingId int, userRequestedId int) *errorHandling.BaseError {
	_, err := repo.Db.Exec(context.Background(), CREATE_FRIEND_REQUEST, userRequestingId, userRequestedId)
	if err != nil {
		return errorHandling.NewBadRequest("failed to send friend request")
	}
	return nil
}

func (repo FriendRequestRepo) RejectFriendRequest(requestId int) *errorHandling.BaseError {
	return nil
}
