package repo

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/hellicopthecat/aisha_do_dot/routines_detail_desc/dtos"
)

func (r *RoutineDetailDescRepo) CreateRoutineDetailDescRepo(ctx context.Context, dto *dtos.CreateRoutineDetailDescDto) error {
	file, err := os.ReadFile("db/ddls/routines_detail_desc/createRoutineDetailDesc.sql")
	if err != nil {
		return nil
	}
	result, err := r.db.ExecContext(ctx, string(file),
		&dto.Id,
		&dto.Init_User_Email,
		&dto.Detail_Desc,
		&dto.Routine_Detail_Id,
	)
	if err != nil {
		return nil
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return nil
	}

	if affected != 1 {
		return fmt.Errorf("세부 디테일의 설명이 생성되지 못했습니다. :: %d", err)
	}

	return nil
}
func (r *RoutineDetailDescRepo) CreateRoutineDetailDescRepoWithTx(tx *sql.Tx, ctx context.Context, dto *dtos.CreateRoutineDetailDescDto) error {
	file, err := os.ReadFile("db/ddls/routines_detail_desc/createRoutineDetailDesc.sql")
	if err != nil {
		return nil
	}
	result, err := tx.ExecContext(ctx, string(file),
		&dto.Id,
		&dto.Init_User_Email,
		&dto.Detail_Desc,
		&dto.Routine_Detail_Id,
	)
	if err != nil {
		return nil
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return nil
	}

	if affected != 1 {
		return fmt.Errorf("세부 디테일의 설명이 생성되지 못했습니다. :: %d", err)
	}

	return nil
}
