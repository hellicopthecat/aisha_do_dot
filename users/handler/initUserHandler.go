package handler

import "github.com/hellicopthecat/aisha_do_dot/users/service"

type UserHandler struct {
	userHandler *service.UserService
}

func InitUserHandler(service service.UserService) UserHandler {
	return UserHandler{userHandler: &service}
}
