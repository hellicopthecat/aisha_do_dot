package handler

import "github.com/hellicopthecat/aisha_do_dot/routines/service"

type RoutineHandler struct {
	routineService *service.RoutineService
}

func InitRoutineHandler(routineService *service.RoutineService) *RoutineHandler {
	return &RoutineHandler{
		routineService: routineService,
	}
}
