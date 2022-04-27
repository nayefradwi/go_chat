package user

import (
	"context"
	"gochat/errorHandling"

	"github.com/jackc/pgx/v4/pgxpool"
)

type IUserRepo interface {
	GetUserById(id int) (User, *errorHandling.BaseError)
	Login(userEmail email, userPassword password) (User, *errorHandling.BaseError)
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

func (repo UserRepo) Login(userEmail email, userPassword password) (User, *errorHandling.BaseError) {
	return User{}, nil
}

func (repo UserRepo) Register(User) (User, *errorHandling.BaseError) {
	return User{}, nil
}
