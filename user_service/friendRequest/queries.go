package friendrequest

const (
	CREATE_FRIEND_REQUEST         = "INSERT INTO friend_requests(user_requesting_id, user_requested_id, created_at) VALUES($1, $2, NOW())"
	GET_FRIEND_REQUESTS           = "select friend_requests.id, username, about, dp from friend_requests inner join users on user_requesting_id=users.id where user_requested_id = $1 and status_id = $2"
	MODIFY_FRIEND_REQUEST         = "update friend_requests set status_id = $1 where user_requested_id = $2 and id = $3 and status_id = $4"
	GET_SENT_FRIEND_REQUESTS      = "select friend_requests.id, username, about, dp from friend_requests inner join users on user_requested_id=users.id where user_requesting_id = $1 and status_id = $2"
	FETCH_USERS_OF_NEW_FRIENDSHIP = "select user1.id, user2.id, user1.username, user2.username, user1.about, user2.about, user1.dp, user2.dp from friend_requests inner join users user1 on user1.id = user_requesting_id inner join users user2 on user2.id = user_requested_id where friend_requests.id = $1 and status_id = $2"
)
