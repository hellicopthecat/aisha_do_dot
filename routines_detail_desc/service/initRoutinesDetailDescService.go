package service

import (
	"github.com/hellicopthecat/aisha_do_dot/routines_detail_desc/repo"
)

type RoutineDetailDescService struct {
	repo *repo.RoutineDetailDescRepo
}

func InitRoutineDetailDesc(repo *repo.RoutineDetailDescRepo) *RoutineDetailDescService {
	return &RoutineDetailDescService{
		repo: repo,
	}
}
