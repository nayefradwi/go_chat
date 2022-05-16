package friendrequest

const (
	StatusPending  int = 0
	StatusAccepted int = 1
	StatusRejected int = 2
)

type FriendRequest struct {
	Id               int `json:"id"`
	UserRequestingId int `json:"userRequestingId"`
	UserRequestedId  int `json:"userRequestedId"`
	Status           int `json:"status"`
}

type FriendRequestDetails struct {
	Id       int    `json:"id"`
	UserId   int    `json:"userId,omitempty"`
	Username string `json:"username"`
	About    string `json:"about,omitempty"`
	Dp       string `json:"dp,omitempty"`
}
