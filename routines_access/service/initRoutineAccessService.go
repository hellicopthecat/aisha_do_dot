package service

import "github.com/hellicopthecat/aisha_do_dot/routines_access/repo"

type RoutineAccessService struct {
	repo *repo.RoutineAccessRepo
}

func InitRoutineAccessService(repo *repo.RoutineAccessRepo) *RoutineAccessService {
	return &RoutineAccessService{
		repo: repo,
	}
}
