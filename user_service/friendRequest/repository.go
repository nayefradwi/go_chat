package friendrequest

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nayefradwi/go_chat/common/errorHandling"
)

type IFriendRequestRepo interface {
	AcceptRequest(ctx context.Context, userId int, requestId int) *errorHandling.BaseError
	GetFriendRequests(ctx context.Context, userId int) ([]FriendRequestDetails, *errorHandling.BaseError)
	SendFriendRequest(ctx context.Context, userRequestingId int, userRequestedId int) *errorHandling.BaseError
	RejectFriendRequest(ctx context.Context, userId int, requestId int) *errorHandling.BaseError
	GetSentFriendRequests(ctx context.Context, userId int) ([]FriendRequestDetails, *errorHandling.BaseError)
}

type FriendRequestRepo struct {
	Db *pgxpool.Pool
}

func (repo FriendRequestRepo) AcceptRequest(ctx context.Context, userId int, requestId int) *errorHandling.BaseError {
	tag, err := repo.Db.Exec(ctx, MODIFY_FRIEND_REQUEST, StatusAccepted, userId, requestId, StatusPending)
	if err != nil || tag.RowsAffected() == 0 {
		return errorHandling.NewBadRequest("failed to accept friend request")
	}
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

func (repo FriendRequestRepo) RejectFriendRequest(ctx context.Context, userId int, requestId int) *errorHandling.BaseError {
	tag, err := repo.Db.Exec(ctx, MODIFY_FRIEND_REQUEST, StatusRejected, userId, requestId, StatusPending)
	if err != nil || tag.RowsAffected() == 0 {
		return errorHandling.NewBadRequest("failed to reject friend request")
	}
	return nil
}

func (repo FriendRequestRepo) GetSentFriendRequests(ctx context.Context, userId int) ([]FriendRequestDetails, *errorHandling.BaseError) {
	friendRequests := make([]FriendRequestDetails, 0)
	rows, err := repo.Db.Query(ctx, GET_SENT_FRIEND_REQUESTS, userId, StatusPending)
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
