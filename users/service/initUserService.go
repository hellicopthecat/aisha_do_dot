package service

import "github.com/hellicopthecat/aisha_do_dot/users/repo"

type UserService struct {
	userService *repo.UserRepo
}

func InitUserService(db repo.UserRepo) UserService {
	return UserService{userService: &db}
}
