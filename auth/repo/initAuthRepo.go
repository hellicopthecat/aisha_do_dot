package repo

import "database/sql"

type AuthRepo struct {
	db *sql.DB
}

func InitAuthRepo(db *sql.DB) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}
