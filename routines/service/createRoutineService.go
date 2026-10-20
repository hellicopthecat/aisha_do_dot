package service

import (
	"context"

	accessDto "github.com/hellicopthecat/aisha_do_dot/routine_access/dtos"
	routineDto "github.com/hellicopthecat/aisha_do_dot/routines/dtos"
)

func (s *RoutineService) CreateRoutineService(ctx context.Context, createRoutineDto routineDto.CreateRoutineDto, createAccessDto accessDto.CreateRoutineAccessDto) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = s.routineRepo.CreateRoutineRepo(ctx, createRoutineDto)
	if err != nil {
		return err
	}

	err = s.accessRepo.CreateRoutineAccessRepo(ctx, createAccessDto)
	if err != nil {
		return err
	}

	return tx.Commit()
}
