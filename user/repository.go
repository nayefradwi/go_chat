package user

type IUserRepo interface {
	GetUserById(id int) (User, error)
	Login(email, password) (User, error)
	Register(User) (User, error)
}

type UserRepo struct {
	// todo: take a db connection pointer
}

func (repo UserRepo) GetUserById(id string) (User, error) { return User{}, nil }

func (repo UserRepo) Login(email string, password string) (User, error) { return User{}, nil }
func (repo UserRepo) Register(email, password) (User, error)            { return User{}, nil }
