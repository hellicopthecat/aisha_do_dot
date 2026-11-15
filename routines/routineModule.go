package routines

import (
	"database/sql"

	aishaJwt "github.com/hellicopthecat/aisha_do_dot/jwt"
	"github.com/hellicopthecat/aisha_do_dot/routines/handler"
	routineRepo "github.com/hellicopthecat/aisha_do_dot/routines/repo"
	"github.com/hellicopthecat/aisha_do_dot/routines/service"
	routineAccessRepo "github.com/hellicopthecat/aisha_do_dot/routines_access/repo"

	echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
)

type RoutineModule struct {
	routineHandler *handler.RoutineHandler
}

func InitRoutineModule(db *sql.DB) *RoutineModule {

	r_repo := routineRepo.InitRoutineRepo(db)
	ra_repo := routineAccessRepo.InitRoutineAccess(db)
	r_service := service.InitRoutineService(db, r_repo, ra_repo)
	r_handler := handler.InitRoutineHandler(r_service)

	return &RoutineModule{
		routineHandler: r_handler,
	}

}

func (h *RoutineModule) RoutineGroupApi(e *echo.Echo) {
	routineGroup := e.Group("/routine")
	config := aishaJwt.EchoJwtConfig()

	routineGroup.GET("", h.routineHandler.CreateRoutineHandler, echojwt.WithConfig(config))

}
