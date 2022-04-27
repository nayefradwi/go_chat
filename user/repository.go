package user

import (
	"context"
	"gochat/auth"
	"gochat/errorHandling"

	"github.com/jackc/pgx/v4/pgxpool"
)

type IUserRepo interface {
	GetUserById(int) (User, *errorHandling.BaseError)
	Login(string, string) (User, *errorHandling.BaseError)
	Register(User) (User, *errorHandling.BaseError)
}

type UserRepo struct {
	Db *pgxpool.Pool
}

func (repo UserRepo) GetUserById(id int) (User, *errorHandling.BaseError) {
	user := User{}
	row := repo.Db.QueryRow(context.Background(), "SELECT id, username, about, dp FROM users WHERE id=$1", id)
	err := row.Scan(&user.Id, &user.Username, &user.About, &user.Dp)
	if err != nil {
		return User{}, errorHandling.NewBadRequest("no user found with this id")
	}
	return user, nil
}

func (repo UserRepo) Login(userEmail string, userPassword string) (User, *errorHandling.BaseError) {
	return User{}, nil
}

func (repo UserRepo) Register(user User) (User, *errorHandling.BaseError) {
	validationErr := checkUserIsValid(user)
	if validationErr != nil {
		return User{}, validationErr
	}
	hash, hashErr := auth.HashPassword(user.Password)
	if hashErr != nil {
		return User{}, errorHandling.NewInternalServerError()
	}
	_, err := repo.Db.Exec(context.Background(), "INSERT INTO users(username, password, email, created_at) VALUES($1, $2, $3, NOW())", user.Username, hash, user.Email)
	if err != nil {
		return User{}, errorHandling.NewBadRequest(err.Error())
	}
	return User{}, nil
}
