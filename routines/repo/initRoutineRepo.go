package repo

import "database/sql"

type RoutineRepo struct {
	db *sql.DB
}

func InitRoutineRepo(db *sql.DB) *RoutineRepo {
	return &RoutineRepo{db: db}
}
