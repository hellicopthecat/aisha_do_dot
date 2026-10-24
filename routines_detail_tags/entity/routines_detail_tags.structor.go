package entity

import (
	"time"

	"github.com/google/uuid"
)

type RoutinesDetailTag struct {
	Id                uuid.UUID
	Created_at        time.Time
	Updated_at        time.Time
	Created_by        string
	Updated_by        string
	Tag               string
	Routine_Detail_Id uuid.UUID
}
