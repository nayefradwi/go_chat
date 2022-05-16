package friendrequest

import (
	"context"
	"encoding/json"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nayefradwi/go_chat/user_service/producer"
	"github.com/nayefradwi/go_chat_common"
)

type IFriendRequestRepo interface {
	AcceptRequest(ctx context.Context, userId int, requestId int) *gochatcommon.BaseError
	GetFriendRequests(ctx context.Context, userId int) ([]FriendRequestDetails, *gochatcommon.BaseError)
	SendFriendRequest(ctx context.Context, userRequestingId int, userRequestedId int) *gochatcommon.BaseError
	RejectFriendRequest(ctx context.Context, userId int, requestId int) *gochatcommon.BaseError
	GetSentFriendRequests(ctx context.Context, userId int) ([]FriendRequestDetails, *gochatcommon.BaseError)
}

type FriendRequestRepo struct {
	Db       *pgxpool.Pool
	Producer producer.IProducer
}

func (repo FriendRequestRepo) AcceptRequest(ctx context.Context, userId int, requestId int) *gochatcommon.BaseError {
	tag, err := repo.Db.Exec(ctx, MODIFY_FRIEND_REQUEST, StatusAccepted, userId, requestId, StatusPending)
	if err != nil || tag.RowsAffected() == 0 {
		return gochatcommon.NewBadRequest("failed to accept friend request")
	}
	go func() {
		row := repo.Db.QueryRow(context.Background(), FETCH_USERS_OF_NEW_FRIENDSHIP, requestId, StatusAccepted)
		repo.produceNewFriendship(row)
	}()
	return nil
}
func (repo FriendRequestRepo) GetFriendRequests(ctx context.Context, userId int) ([]FriendRequestDetails, *gochatcommon.BaseError) {
	friendRequests := make([]FriendRequestDetails, 0)
	rows, err := repo.Db.Query(ctx, GET_FRIEND_REQUESTS, userId, StatusPending)
	if err != nil {
		return friendRequests, gochatcommon.NewBadRequest("failed to load friend requests")
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

func (repo FriendRequestRepo) SendFriendRequest(ctx context.Context, userRequestingId int, userRequestedId int) *gochatcommon.BaseError {
	_, err := repo.Db.Exec(ctx, CREATE_FRIEND_REQUEST, userRequestingId, userRequestedId)
	if err != nil {
		return gochatcommon.NewBadRequest("failed to send friend request")
	}
	return nil
}

func (repo FriendRequestRepo) RejectFriendRequest(ctx context.Context, userId int, requestId int) *gochatcommon.BaseError {
	tag, err := repo.Db.Exec(ctx, MODIFY_FRIEND_REQUEST, StatusRejected, userId, requestId, StatusPending)
	if err != nil || tag.RowsAffected() == 0 {
		return gochatcommon.NewBadRequest("failed to reject friend request")
	}
	return nil
}

func (repo FriendRequestRepo) GetSentFriendRequests(ctx context.Context, userId int) ([]FriendRequestDetails, *gochatcommon.BaseError) {
	friendRequests := make([]FriendRequestDetails, 0)
	rows, err := repo.Db.Query(ctx, GET_SENT_FRIEND_REQUESTS, userId, StatusPending)
	if err != nil {
		return friendRequests, gochatcommon.NewBadRequest("failed to load friend requests")
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

func (repo FriendRequestRepo) produceNewFriendship(row pgx.Row) {
	user1, user2, err := generateNewFriendShipUsersFromQueryRow(row)
	if err != nil {
		return
	}
	usersMap := make(map[string]FriendRequestDetails)
	usersMap["user1"] = user1
	usersMap["user2"] = user2
	event, _ := json.Marshal(usersMap)
	repo.Producer.CreateJsonEvent(producer.NewFriendshipTopic, event)
}

func generateNewFriendShipUsersFromQueryRow(row pgx.Row) (FriendRequestDetails, FriendRequestDetails, error) {
	var user1Id, user2Id int
	var user1Username, user2Username, user1About, user2About, user1Dp, user2Dp string
	err := row.Scan(&user1Id, &user2Id, &user1Username, &user2Username, &user1About, &user2About, &user1Dp, &user2Dp)
	if err != nil {
		log.Printf("error scanning new friendship: %s", err.Error())
		return FriendRequestDetails{}, FriendRequestDetails{}, err
	}
	user1, user2 := FriendRequestDetails{
		UserId:   user1Id,
		Username: user1Username,
		About:    user1About,
		Dp:       user1Dp,
	}, FriendRequestDetails{
		UserId:   user2Id,
		Username: user2Username,
		About:    user2About,
		Dp:       user2Dp,
	}
	return user1, user2, nil
}
