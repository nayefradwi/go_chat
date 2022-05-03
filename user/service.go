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
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		goChatUtil.WriteErrorResponse(w, errorHandling.NewInternalServerError())
		return
	}
	authUser, loginErr := service.userRepo.Login(user.Email, user.Password)
	if loginErr != nil {
		goChatUtil.WriteErrorResponse(w, loginErr)
		return
	}
	json.NewEncoder(w).Encode(authUser)
}

func (service UserService) Register(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		goChatUtil.WriteErrorResponse(w, errorHandling.NewInternalServerError())
		return
	}
	registrationErr := service.userRepo.Register(user)
	if registrationErr != nil {
		goChatUtil.WriteErrorResponse(w, registrationErr)
		return
	}
	json.NewEncoder(w).Encode(make(map[string]string))
}

func (service UserService) GetUserById(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(auth.UserIdKey{}).(int)
	user, err := service.userRepo.GetUserById(userId)
	if err != nil {
		goChatUtil.WriteErrorResponse(w, err)
		return
	}
	json.NewEncoder(w).Encode(user)
}
