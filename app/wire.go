//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/labstack/echo"
	http2 "goWebServer/internal/users/delivery/http"
	gorm2 "goWebServer/internal/users/repository/mySql"
	"goWebServer/internal/users/usecase"
	"gorm.io/gorm"
	"time"
)

var UserDependencySet = wire.NewSet(
	gorm2.NewUserPersistenceRepository,
	usecase.NewUserUseCase,
	http2.NewUserCommandHandler,
	http2.NewUserQueryHandler,
)

func InitializeWebServerDependency(db *gorm.DB, echoInstance *echo.Echo, timeout time.Duration) *WebServer {
	wire.Build(
		NewUserServer,
		UserDependencySet)
	return nil
}
