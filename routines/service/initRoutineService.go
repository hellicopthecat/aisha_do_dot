package service

import (
	"database/sql"

	accessRepo "github.com/hellicopthecat/aisha_do_dot/routine_access/repo"
	routineRepo "github.com/hellicopthecat/aisha_do_dot/routines/repo"
)

type RoutineService struct {
	db          *sql.DB
	routineRepo *routineRepo.RoutineRepo
	accessRepo  *accessRepo.RoutineAccessRepo
}

func InitRoutineService(db *sql.DB, routineRepo *routineRepo.RoutineRepo, accessRepo *accessRepo.RoutineAccessRepo) *RoutineService {
	return &RoutineService{
		db:          db,
		routineRepo: routineRepo,
		accessRepo:  accessRepo,
	}
}
