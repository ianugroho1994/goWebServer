package handler

import (
	"hardtmann/api/authentication/usecases"

	"github.com/labstack/echo/v4"
)

type AuthenticationAPI struct {
	commands *AuthenticationAPICommands
}

func NewAuthenticationAPI() *AuthenticationAPI {
	useCase := usecases.NewAuthenticationUseCase()
	return &AuthenticationAPI{
		commands: NewAuthenticationAPICommand(useCase),
	}
}

func (a *AuthenticationAPI) PopulateRouteHandler(e *echo.Echo) {
	group := e.Group("auth")
	a.commands.PopulateCommandAPI(group)
}
