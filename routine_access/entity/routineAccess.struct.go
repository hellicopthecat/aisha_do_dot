package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/hellicopthecat/aisha_do_dot/commons/enums"
)

type RoutineAccessEntity struct {
	Id             uuid.UUID
	created_at     time.Time
	updated_at     time.Time
	created_by     string
	updated_by     string
	can_access     enums.Access
	editable       enums.AccessRole
	routine_id     uuid.UUID
	access_user_id uuid.UUID
}
