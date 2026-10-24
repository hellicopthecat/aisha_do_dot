package service

import (
	"context"

	routineDto "github.com/hellicopthecat/aisha_do_dot/routines/dtos"
	accessDto "github.com/hellicopthecat/aisha_do_dot/routines_access/dtos"
)

func (s *RoutineService) CreateRoutineService(ctx context.Context, createRoutineDto routineDto.CreateRoutineDto, createAccessDto accessDto.CreateRoutineAccessDto) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = s.routineRepo.CreateRoutineRepo(ctx, tx, createRoutineDto)
	if err != nil {
		return err
	}

	err = s.accessRepo.CreateRoutineAccessRepoWithTx(ctx, tx, createAccessDto)
	if err != nil {
		return err
	}

	return tx.Commit()
}
