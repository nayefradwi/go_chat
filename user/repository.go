package user

import (
	"context"
	"gochat/auth"
	"gochat/errorHandling"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type IUserRepo interface {
	GetUserById(int) (User, *errorHandling.BaseError)
	Login(string, string) (AuthenticatedUser, *errorHandling.BaseError)
	Register(User) *errorHandling.BaseError
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

func (repo UserRepo) Login(userEmail string, userPassword string) (AuthenticatedUser, *errorHandling.BaseError) {
	log.Printf("email: %s, password: %s", userEmail, userPassword)
	user := User{}
	row := repo.Db.QueryRow(context.Background(), "SELECT id, username, about, dp, password FROM users WHERE email=$1", userEmail)
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

func (repo UserRepo) Register(user User) *errorHandling.BaseError {
	validationErr := checkUserIsValid(user)
	if validationErr != nil {
		return validationErr
	}
	hash, hashErr := auth.HashPassword(user.Password)
	if hashErr != nil {
		return errorHandling.NewInternalServerError()
	}
	_, err := repo.Db.Exec(context.Background(), "INSERT INTO users(username, password, email, created_at) VALUES($1, $2, $3, NOW())", user.Username, hash, user.Email)
	if err != nil {
		return errorHandling.NewBadRequest(err.Error())
	}
	return nil
}
