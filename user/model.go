package user

type email string
type password string

type User struct {
	Id       int      `json:"id,omitempty"`
	Username string   `json:"username,omitempty"`
	About    string   `json:"about,omitempty"`
	Dp       string   `json:"dp,omitempty"`
	Email    email    `json:"email,omitempty"`
	Password password `json:"password,omitempty"`
}
