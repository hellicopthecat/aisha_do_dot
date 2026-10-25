package service

import (
	"context"

	"github.com/hellicopthecat/aisha_do_dot/routines_detail_desc/dtos"
)

func (s *RoutineDetailDescService) CreateRoutinesDetailDescService(ctx context.Context, dto *dtos.CreateRoutineDetailDescDto) error {
	return s.repo.CreateRoutineDetailDescRepo(ctx, dto)
}
