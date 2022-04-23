package user

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"gochat/errorHandling"
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
	return User{}, nil
}

func (repo UserRepo) Login(userEmail email, userPassword password) (User, *errorHandling.BaseError) {
	return User{}, nil
}
func (repo UserRepo) Register(User) (User, *errorHandling.BaseError) {
	return User{}, nil
}
