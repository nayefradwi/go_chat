package user

type User struct {
	Id       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	About    string `json:"about,omitempty"`
	Dp       string `json:"dp,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type AuthenticatedUser struct {
	AccessToken string `json:"accessToken"`
}
