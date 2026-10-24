package service

import (
	"database/sql"

	routineRepo "github.com/hellicopthecat/aisha_do_dot/routines/repo"
	accessRepo "github.com/hellicopthecat/aisha_do_dot/routines_access/repo"
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
