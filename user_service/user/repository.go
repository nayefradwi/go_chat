package user

import (
	"context"
	"github.com/nayefradwi/go_chat/common/errorHandling"
	"github.com/nayefradwi/go_chat/user_service/auth"

	"github.com/jackc/pgx/v4/pgxpool"
)

type IUserRepo interface {
	GetUserById(context.Context, int) (User, *errorHandling.BaseError)
	Login(context.Context, string, string) (AuthenticatedUser, *errorHandling.BaseError)
	Register(context.Context, User) *errorHandling.BaseError
}

type UserRepo struct {
	Db *pgxpool.Pool
}

func (repo UserRepo) GetUserById(ctx context.Context, id int) (User, *errorHandling.BaseError) {
	user := User{}
	row := repo.Db.QueryRow(ctx, "SELECT id, username, about, dp FROM users WHERE id=$1", id)
	err := row.Scan(&user.Id, &user.Username, &user.About, &user.Dp)
	if err != nil {
		return User{}, errorHandling.NewBadRequest("no user found with this id")
	}
	return user, nil
}

func (repo UserRepo) Login(ctx context.Context, userEmail string, userPassword string) (AuthenticatedUser, *errorHandling.BaseError) {
	user := User{}
	row := repo.Db.QueryRow(ctx, "SELECT id, username, about, dp, password FROM users WHERE email=$1", userEmail)
	err := row.Scan(&user.Id, &user.Username, &user.About, &user.Dp, &user.Password)
	if err != nil {
		return AuthenticatedUser{}, errorHandling.NewBadRequest(err.Error())
	}
	if isCorrectPassword := auth.ComparePassword(userPassword, user.Password); !isCorrectPassword {
		return AuthenticatedUser{}, errorHandling.NewBadRequest("password is incorrect")
	}
	token, jwtErr := auth.GenerateToken(user.Id)
	if jwtErr != nil {
		return AuthenticatedUser{}, errorHandling.NewInternalServerError()
	}
	authUser := AuthenticatedUser{
		AccessToken: token,
	}
	return authUser, nil
}

func (repo UserRepo) Register(ctx context.Context, user User) *errorHandling.BaseError {
	validationErr := checkUserIsValid(user)
	if validationErr != nil {
		return validationErr
	}
	hash, hashErr := auth.HashPassword(user.Password)
	if hashErr != nil {
		return errorHandling.NewInternalServerError()
	}
	_, err := repo.Db.Exec(ctx, "INSERT INTO users(username, password, email, created_at, dp, about) VALUES($1, $2, $3, NOW(), $4, $5)", user.Username, hash, user.Email, user.Dp, user.About)
	if err != nil {
		return errorHandling.NewBadRequest(err.Error())
	}
	return nil
}
