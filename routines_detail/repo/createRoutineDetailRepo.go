package repo

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/hellicopthecat/aisha_do_dot/routines_detail/dtos"
)

func (r *RoutineDetailRepo) CreateRoutineDetailRepo(ctx context.Context, dto *dtos.CreateRoutineDetailDto) error {

	file, err := os.ReadFile("db/ddls/routines_detail/createRoutinesDetail.sql")
	if err != nil {
		return err
	}
	result, err := r.db.ExecContext(ctx, string(file),
		&dto.Id,
		&dto.Init_User_Email,
		&dto.Init_User_Email,
		&dto.Priority_Two_Depth,
		&dto.Pre_Event_Start_At,
		&dto.Pre_Event_End_At,
		&dto.Start_At,
		&dto.End_At,
		&dto.Routine_Id,
	)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected != 1 {
		return fmt.Errorf("세부루틴이 생성되지 않았습니다. :: %d", err)
	}
	return nil
}

func (r *RoutineDetailRepo) CreateRoutineDetailRepoWithTx(tx *sql.Tx, ctx context.Context, dto *dtos.CreateRoutineDetailDto) error {
	file, err := os.ReadFile("db/ddls/routines_detail/createRoutinesDetail.sql")
	if err != nil {
		return err
	}
	result, err := tx.ExecContext(ctx, string(file),
		&dto.Id,
		&dto.Init_User_Email,
		&dto.Init_User_Email,
		&dto.Priority_Two_Depth,
		&dto.Pre_Event_Start_At,
		&dto.Pre_Event_End_At,
		&dto.Start_At,
		&dto.End_At,
		&dto.Routine_Id,
	)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected != 1 {
		return fmt.Errorf("세부루틴이 생성되지 않았습니다. :: %d", err)
	}
	return nil
}
