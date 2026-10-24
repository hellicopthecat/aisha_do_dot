package repo

import (
	"context"
	"database/sql"
	"os"

	"github.com/hellicopthecat/aisha_do_dot/routines_detail_tags/dtos"
)

func (r *RoutineDetailTagsRepo) CreateRoutineDetailTagsRepo(ctx context.Context, dto *dtos.CreateRoutineDetailTagsDto) error {
	file, err := os.ReadFile("db/ddls/routines_detail_tags/createRoutinesDetailTags.sql")
	if err != nil {
		return err
	}
	result, err := r.db.ExecContext(ctx, string(file),
		&dto.Id,
		&dto.Init_User_Email,
		&dto.Tag,
		&dto.Routine_Detail_Id,
	)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected != 1 {
		return err
	}
	return nil
}
func (r *RoutineDetailTagsRepo) CreateRoutineDetailTagsRepoWithTx(tx *sql.Tx, ctx context.Context, dto *dtos.CreateRoutineDetailTagsDto) error {
	file, err := os.ReadFile("db/ddls/routines_detail_tags/createRoutinesDetailTags.sql")
	if err != nil {
		return err
	}
	result, err := tx.ExecContext(ctx, string(file),
		&dto.Id,
		&dto.Init_User_Email,
		&dto.Tag,
		&dto.Routine_Detail_Id,
	)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected != 1 {
		return err
	}
	return nil
}
