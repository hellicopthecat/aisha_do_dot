package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/hellicopthecat/aisha_do_dot/routines_detail/dtos"
)

func (s *RoutineDetailService) FindRoutinesDetailByRoutinesIdService(ctx context.Context, routine_id uuid.UUID) (*[]dtos.FindRoutineDetailByRoutineIdDto, error) {
	return s.routineDetailRepo.FindRoutineDetailByRoutineIdRepo(ctx, routine_id)
}
