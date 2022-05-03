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
