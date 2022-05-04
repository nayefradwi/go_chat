package user

import (
	"encoding/json"
	"gochat/auth"
	"gochat/errorHandling"
	"gochat/goChatUtil"
	"net/http"
)

type UserService struct {
	userRepo IUserRepo
}

func NewUserService(userRepo IUserRepo) UserService {
	return UserService{
		userRepo: userRepo,
	}
}

func (service UserService) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		goChatUtil.WriteErrorResponse(w, errorHandling.NewInternalServerError())
		return
	}
	authUser, loginErr := service.userRepo.Login(ctx, user.Email, user.Password)
	if loginErr != nil {
		goChatUtil.WriteErrorResponse(w, loginErr)
		return
	}
	json.NewEncoder(w).Encode(authUser)
}

func (service UserService) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		goChatUtil.WriteErrorResponse(w, errorHandling.NewInternalServerError())
		return
	}
	registrationErr := service.userRepo.Register(ctx, user)
	if registrationErr != nil {
		goChatUtil.WriteErrorResponse(w, registrationErr)
		return
	}
	goChatUtil.WriteEmptyCreatedResponse(w)
}

func (service UserService) GetUserById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value(auth.UserIdKey{}).(int)
	user, err := service.userRepo.GetUserById(ctx, userId)
	if err != nil {
		goChatUtil.WriteErrorResponse(w, err)
		return
	}
	json.NewEncoder(w).Encode(user)
}
