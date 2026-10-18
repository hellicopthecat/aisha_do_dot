package repo

import (
	"context"
	"database/sql"
	"errors"
	"os"
)

func (r UserRepo) FindUserByEmailRepo(ctx context.Context, email string) (*string, error) {
	file, err := os.ReadFile("db/ddls/users/findUserByEmail.sql")
	if err != nil {
		return nil, err
	}

	var existEmail string

	err = r.db.QueryRowContext(ctx, string(file), email).Scan(&existEmail)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &existEmail, nil
}
