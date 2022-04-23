package user

import "net/http"

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

}

func (service UserService) GetUserById(w http.ResponseWriter, r *http.Request) {
	_, err := service.userRepo.GetUserById(1)
	if err != nil {
		response := err.GenerateResponse()
		w.WriteHeader(err.Status)
		w.Write(response)
	}
}
