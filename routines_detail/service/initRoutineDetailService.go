package service

import (
	"database/sql"

	rd_repo "github.com/hellicopthecat/aisha_do_dot/routines_detail/repo"
	rdd_repo "github.com/hellicopthecat/aisha_do_dot/routines_detail_desc/repo"
	rdt_repo "github.com/hellicopthecat/aisha_do_dot/routines_detail_tags/repo"
)

type RoutineDetailService struct {
	db                    *sql.DB
	routineDetailRepo     *rd_repo.RoutineDetailRepo
	routineDetailDescRepo *rdd_repo.RoutineDetailDescRepo
	routineDetailTagsRepo *rdt_repo.RoutineDetailTagsRepo
}

func InitRoutineDetailService(db *sql.DB,
	routineDetailRepo *rd_repo.RoutineDetailRepo,
	routineDetailDescRepo *rdd_repo.RoutineDetailDescRepo,
) *RoutineDetailService {
	return &RoutineDetailService{
		db:                    db,
		routineDetailRepo:     routineDetailRepo,
		routineDetailDescRepo: routineDetailDescRepo,
	}
}
