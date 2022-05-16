package user

import (
	"encoding/json"
	"github.com/nayefradwi/go_chat/user_service/middleware"
	"github.com/nayefradwi/go_chat_common"
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
		gochatcommon.WriteErrorResponse(w, gochatcommon.NewInternalServerError())
		return
	}
	authUser, loginErr := service.userRepo.Login(ctx, user.Email, user.Password)
	if loginErr != nil {
		gochatcommon.WriteErrorResponse(w, loginErr)
		return
	}
	json.NewEncoder(w).Encode(authUser)
}

func (service UserService) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		gochatcommon.WriteErrorResponse(w, gochatcommon.NewInternalServerError())
		return
	}
	registrationErr := service.userRepo.Register(ctx, user)
	if registrationErr != nil {
		gochatcommon.WriteErrorResponse(w, registrationErr)
		return
	}
	gochatcommon.WriteEmptyCreatedResponse(w, "user registered")
}

func (service UserService) GetUserById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value(middleware.UserIdKey{}).(int)
	user, err := service.userRepo.GetUserById(ctx, userId)
	if err != nil {
		gochatcommon.WriteErrorResponse(w, err)
		return
	}
	json.NewEncoder(w).Encode(user)
}
