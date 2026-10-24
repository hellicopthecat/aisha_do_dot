package service

import (
	"context"

	rd_dto "github.com/hellicopthecat/aisha_do_dot/routines_detail/dtos"
	rdd_dto "github.com/hellicopthecat/aisha_do_dot/routines_detail_desc/dtos"
	rdt_dto "github.com/hellicopthecat/aisha_do_dot/routines_detail_tags/dtos"
)

func (s *RoutineDetailService) CreateRoutineDetailServiceWithTx(ctx context.Context, rdDto *rd_dto.CreateRoutineDetailDto, rddDto *rdd_dto.CreateRoutineDetailDescDto, rdtDto *rdt_dto.CreateRoutineDetailTagsDto) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = s.routineDetailRepo.CreateRoutineDetailRepoWithTx(tx, ctx, rdDto)
	if err != nil {
		return err
	}
	err = s.routineDetailDescRepo.CreateRoutineDetailDescRepoWithTx(tx, ctx, rddDto)
	if err != nil {
		return err
	}

	if rdtDto != nil {
		err = s.routineDetailTagsRepo.CreateRoutineDetailTagsRepoWithTx(tx, ctx, rdtDto)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
