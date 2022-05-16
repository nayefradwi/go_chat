package user

import (
	"context"
	"encoding/json"

	"github.com/nayefradwi/go_chat/chat_service/config"
	"github.com/nayefradwi/go_chat/common/auth"
	"github.com/nayefradwi/go_chat/common/errorHandling"
	"github.com/nayefradwi/go_chat/user_service/producer"

	"github.com/jackc/pgx/v4/pgxpool"
)

type IUserRepo interface {
	GetUserById(context.Context, int) (User, *errorHandling.BaseError)
	Login(context.Context, string, string) (AuthenticatedUser, *errorHandling.BaseError)
	Register(context.Context, User) *errorHandling.BaseError
}

type UserRepo struct {
	Db       *pgxpool.Pool
	Producer producer.IProducer
}

func (repo UserRepo) GetUserById(ctx context.Context, id int) (User, *errorHandling.BaseError) {
	user := User{}
	row := repo.Db.QueryRow(ctx, GET_USER_BY_ID, id)
	err := row.Scan(&user.Id, &user.Username, &user.About, &user.Dp)
	if err != nil {
		return User{}, errorHandling.NewBadRequest("no user found with this id")
	}
	return user, nil
}

func (repo UserRepo) Login(ctx context.Context, userEmail string, userPassword string) (AuthenticatedUser, *errorHandling.BaseError) {
	user := User{}
	row := repo.Db.QueryRow(ctx, LOGIN, userEmail)
	err := row.Scan(&user.Id, &user.Username, &user.About, &user.Dp, &user.Password)
	if err != nil {
		return AuthenticatedUser{}, errorHandling.NewBadRequest(err.Error())
	}
	if isCorrectPassword := auth.ComparePassword(userPassword, user.Password); !isCorrectPassword {
		return AuthenticatedUser{}, errorHandling.NewBadRequest("password is incorrect")
	}
	token, jwtErr := auth.GenerateToken(user.Id, config.Secret)
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
	_, err := repo.Db.Exec(ctx, REGISTER, user.Username, hash, user.Email, user.Dp, user.About)
	if err != nil {
		return errorHandling.NewBadRequest(err.Error())
	}
	data, _ := json.Marshal(user)
	go repo.Producer.CreateJsonEvent(producer.UserRegisteredTopic, data)
	return nil
}
