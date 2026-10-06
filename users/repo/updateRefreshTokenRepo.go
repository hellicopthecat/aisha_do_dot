package repo

import (
	"context"
	"fmt"
	"os"

	dtos "github.com/hellicopthecat/aisha_do_dot/users/dtos/request"
)

func (r UserRepo) UpdateRefreshTokenRepo(ctx context.Context, dto dtos.UpdateRefreshToken) error {
	file, err := os.ReadFile("db/ddls/user/update_refresh_token.sql")
	if err != nil {
		return err
	}

	result, err := r.db.ExecContext(ctx, string(file), dto.Refresh_token, dto.Id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return fmt.Errorf("리프레시 토큰 저장 실패 :: %v", err)
	}

	return nil
}
