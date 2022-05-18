package user

import (
	"context"

	"github.com/nayefradwi/go_chat/user_service/config"
	"github.com/nayefradwi/go_chat/user_service/producer"
	"github.com/nayefradwi/go_chat_common"

	"github.com/jackc/pgx/v4/pgxpool"
)

type IUserRepo interface {
	GetUserById(context.Context, int) (User, *gochatcommon.BaseError)
	Login(context.Context, string, string) (AuthenticatedUser, *gochatcommon.BaseError)
	Register(context.Context, User) *gochatcommon.BaseError
}

type UserRepo struct {
	Db       *pgxpool.Pool
	Producer producer.IProducer
}

func (repo UserRepo) GetUserById(ctx context.Context, id int) (User, *gochatcommon.BaseError) {
	user := User{}
	row := repo.Db.QueryRow(ctx, GET_USER_BY_ID, id)
	err := row.Scan(&user.Id, &user.Username, &user.About, &user.Dp)
	if err != nil {
		return User{}, gochatcommon.NewBadRequest("no user found with this id")
	}
	return user, nil
}

func (repo UserRepo) Login(ctx context.Context, userEmail string, userPassword string) (AuthenticatedUser, *gochatcommon.BaseError) {
	user := User{}
	row := repo.Db.QueryRow(ctx, LOGIN, userEmail)
	err := row.Scan(&user.Id, &user.Username, &user.About, &user.Dp, &user.Password)
	if err != nil {
		return AuthenticatedUser{}, gochatcommon.NewBadRequest(err.Error())
	}
	if isCorrectPassword := gochatcommon.ComparePassword(userPassword, user.Password); !isCorrectPassword {
		return AuthenticatedUser{}, gochatcommon.NewBadRequest("password is incorrect")
	}
	token, jwtErr := gochatcommon.GenerateToken(user.Id, config.Secret)
	if jwtErr != nil {
		return AuthenticatedUser{}, gochatcommon.NewInternalServerError()
	}
	authUser := AuthenticatedUser{
		AccessToken: token,
	}
	return authUser, nil
}

func (repo UserRepo) Register(ctx context.Context, user User) *gochatcommon.BaseError {
	validationErr := checkUserIsValid(user)
	if validationErr != nil {
		return validationErr
	}
	hash, hashErr := gochatcommon.HashPassword(user.Password)
	if hashErr != nil {
		return gochatcommon.NewInternalServerError()
	}
	_, err := repo.Db.Exec(ctx, REGISTER, user.Username, hash, user.Email, user.Dp, user.About)
	if err != nil {
		return gochatcommon.NewBadRequest(err.Error())
	}
	// data, _ := json.Marshal(user)
	// go repo.Producer.CreateJsonEvent(producer.UserRegisteredTopic, data)
	return nil
}
