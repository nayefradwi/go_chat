package user

const (
	GET_USER_BY_ID = "SELECT id, username, about, dp FROM users WHERE id=$1"
	LOGIN          = "SELECT id, username, about, dp, password FROM users WHERE email=$1"
	REGISTER       = "INSERT INTO users(username, password, email, created_at, dp, about) VALUES($1, $2, $3, NOW(), $4, $5)"
)
