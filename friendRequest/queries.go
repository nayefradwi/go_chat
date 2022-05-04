package friendrequest

const (
	CREATE_FRIEND_REQUEST = "INSERT INTO friend_requests(user_requesting_id, user_requested_id, created_at) VALUES($1, $2, NOW())"
	GET_FRIEND_REQUESTS   = "select friend_requests.id, username, about, dp from friend_requests inner join users on user_requested_id=users.id where users.id = $1 and status_id = $2"
)
