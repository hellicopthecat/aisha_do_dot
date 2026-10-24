package repo

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/hellicopthecat/aisha_do_dot/routines_access/dtos"
)

func (r *RoutineAccessRepo) CreateRoutineAccessRepo(ctx context.Context, dto dtos.CreateRoutineAccessDto) error {
	file, err := os.ReadFile("db/ddls/routine_access/createRoutineAccess.sql")
	if err != nil {
		return err
	}

	result, err := r.db.ExecContext(ctx, string(file),
		dto.Id,
		dto.User_Email,
		dto.User_Email,
		dto.Access,
		dto.Editable,
		dto.Routine_Id,
		dto.Access_User_Id,
	)

	if err != nil {
		return err
	}

	affect, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if affect != 1 {
		return fmt.Errorf("생성되지 않았습니다. %d", affect)
	}

	return nil
}

func (r *RoutineAccessRepo) CreateRoutineAccessRepoWithTx(ctx context.Context, tx *sql.Tx, dto dtos.CreateRoutineAccessDto) error {
	file, err := os.ReadFile("db/ddls/routine_access/createRoutineAccess.sql")
	if err != nil {
		return err
	}

	result, err := tx.ExecContext(ctx, string(file),
		dto.Id,
		dto.User_Email,
		dto.User_Email,
		dto.Access,
		dto.Editable,
		dto.Routine_Id,
		dto.Access_User_Id,
	)

	if err != nil {
		return err
	}

	affect, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if affect != 1 {
		return fmt.Errorf("생성되지 않았습니다. %d", affect)
	}

	return nil
}
