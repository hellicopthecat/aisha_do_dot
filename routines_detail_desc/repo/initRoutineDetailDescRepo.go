package repo

import "database/sql"

type RoutineDetailDescRepo struct {
	db *sql.DB
}

func InitRoutineDetailRepo(db *sql.DB) *RoutineDetailDescRepo {
	return &RoutineDetailDescRepo{
		db: db,
	}
}
