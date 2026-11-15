package repo

import "database/sql"

type RoutineAccessRepo struct {
	db *sql.DB
}

func InitRoutineAccess(db *sql.DB) *RoutineAccessRepo {
	return &RoutineAccessRepo{db: db}
}
