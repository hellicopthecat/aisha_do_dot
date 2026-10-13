package repo

import (
	"context"
	"os"

	"github.com/google/uuid"
)

func (r UserRepo) GetUserIdByEmailRepo(ctx context.Context, email string) (*uuid.UUID, error) {
	file, err := os.ReadFile("db/ddls/user/getUserIdEmail.sql")
	if err != nil {
		return nil, err
	}

	var uid uuid.UUID
	err = r.db.QueryRowContext(ctx, string(file), email).Scan(&uid)
	if err != nil {
		return nil, err
	}

	return &uid, nil
}
