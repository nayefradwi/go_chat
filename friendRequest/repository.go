package friendrequest

import (
	"context"
	"gochat/errorHandling"

	"github.com/jackc/pgx/v4/pgxpool"
)

type IFriendRequestRepo interface {
	AcceptRequest(ctx context.Context, requestId int) *errorHandling.BaseError
	GetFriendRequests(ctx context.Context, userId int) ([]FriendRequestDetails, *errorHandling.BaseError)
	SendFriendRequest(ctx context.Context, userRequestingId int, userRequestedId int) *errorHandling.BaseError
	RejectFriendRequest(ctx context.Context, requestId int) *errorHandling.BaseError
}

type FriendRequestRepo struct {
	Db *pgxpool.Pool
}

func (repo FriendRequestRepo) AcceptRequest(ctx context.Context, requestId int) *errorHandling.BaseError {
	repo.Db.Exec(ctx, "")
	return nil
}
func (repo FriendRequestRepo) GetFriendRequests(ctx context.Context, userId int) ([]FriendRequestDetails, *errorHandling.BaseError) {
	friendRequests := make([]FriendRequestDetails, 0)
	rows, err := repo.Db.Query(ctx, GET_FRIEND_REQUESTS, userId, StatusPending)
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

func (repo FriendRequestRepo) SendFriendRequest(ctx context.Context, userRequestingId int, userRequestedId int) *errorHandling.BaseError {
	_, err := repo.Db.Exec(ctx, CREATE_FRIEND_REQUEST, userRequestingId, userRequestedId)
	if err != nil {
		return errorHandling.NewBadRequest("failed to send friend request")
	}
	return nil
}

func (repo FriendRequestRepo) RejectFriendRequest(ctx context.Context, requestId int) *errorHandling.BaseError {
	return nil
}
