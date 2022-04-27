package user

import (
	"encoding/json"
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

}
func (service UserService) Register(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		goChatUtil.WriteErrorResponse(w, errorHandling.NewInternalServerError())
		return
	}
	user, registrationErr := service.userRepo.Register(user)
	if registrationErr != nil {
		goChatUtil.WriteErrorResponse(w, registrationErr)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (service UserService) GetUserById(w http.ResponseWriter, r *http.Request) {
	user, err := service.userRepo.GetUserById(1)
	if err != nil {
		goChatUtil.WriteErrorResponse(w, err)
		return
	}
	json.NewEncoder(w).Encode(user)
}
