package repo

import "database/sql"

type RoutineDetailRepo struct {
	db *sql.DB
}

func InitRoutineDetailRepo(db *sql.DB) *RoutineDetailRepo {
	return &RoutineDetailRepo{
		db: db,
	}
}
