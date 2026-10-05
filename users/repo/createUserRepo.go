package repo

import (
	"context"
	"fmt"
	"os"

	"github.com/hellicopthecat/aisha_do_dot/users/entity"
)

func (r UserRepo) CreateUserRepo(ctx context.Context, userEntity entity.User) error {
	file, err := os.ReadFile("db/ddls/user/create_user.sql")
	if err != nil {
		return err
	}

	result, err := r.db.ExecContext(ctx, string(file),
		userEntity.Id,            // id
		userEntity.Email,         // created by
		userEntity.Email,         // updated by
		userEntity.Email,         // emial
		userEntity.Social,        // social
		userEntity.Provider_id,   // provider_id
		userEntity.Name,          // name
		userEntity.Refresh_token, // refresh_token
	)

	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return fmt.Errorf("유저 생성 실패 :: %d", rows)
	}
	return nil
}
