package handler

import (
	"hardtmann/smartlab/api/user/domain"
	"hardtmann/smartlab/api/user/usecases"
	"time"

	"github.com/labstack/echo/v4"
)

type UserAPI struct {
	commands *UserAPICommands
	queries  *UserAPIQueries
}

func NewUserAPI(userRepo domain.IUserRepository, timeout time.Duration) *UserAPI {
	useCase := usecases.NewUserUseCase(userRepo, timeout)
	return &UserAPI{
		commands: NewUserAPICommand(useCase, userRepo),
		queries:  NewUserAPIQueries(userRepo),
	}
}

func (o *UserAPI) PopulateRouteHandler(e *echo.Echo) {
	route := e.Group("/user")
	o.commands.PopulateCommandAPI(route)
	o.queries.PopulateQueriesAPI(route)
}
