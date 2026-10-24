package repo

import "database/sql"

type RoutineDetailTagsRepo struct {
	db *sql.DB
}

func InitRoutineDetailTags(db *sql.DB) *RoutineDetailTagsRepo {
	return &RoutineDetailTagsRepo{db: db}
}
