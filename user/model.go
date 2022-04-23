package user

type email string
type password string

type User struct {
	Id       int      `json:"id"`
	Username string   `json:"username"`
	About    string   `json:"about"`
	Dp       string   `json:"dp"`
	Email    email    `json:"email"`
	Password password `json:"password"`
}

func (password) MarshalJSON() ([]byte, error) {
	return []byte(""), nil
}

func (email) MarshalJSON() ([]byte, error) {
	return []byte(""), nil
}
