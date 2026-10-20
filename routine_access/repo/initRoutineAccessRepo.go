package repo

import "database/sql"

type RoutineAccessRepo struct {
	db *sql.DB
}

func (r RoutineAccessRepo) InitRoutineAccess(db *sql.DB) *RoutineAccessRepo {
	return &RoutineAccessRepo{db: db}
}
