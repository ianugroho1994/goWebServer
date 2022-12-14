// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/google/wire"
	"github.com/labstack/echo"
	"goWebServer/internal/users/delivery/http"
	"goWebServer/internal/users/repository/mySql"
	"goWebServer/internal/users/usecase"
	"gorm.io/gorm"
	"time"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func InitializeWebServerDependency(db *gorm.DB, echoInstance *echo.Echo, timeout time.Duration) *WebServer {
	userPersistentRepository := mySql.NewUserPersistenceRepository(db)
	userUseCase := usecase.NewUserUseCase(userPersistentRepository, timeout)
	userCommandHandler := http.NewUserCommandHandler(echoInstance, userUseCase)
	userQueryHandler := http.NewUserQueryHandler(echoInstance, userUseCase)
	webServer := NewUserServer(echoInstance, userCommandHandler, userQueryHandler)
	return webServer
}

// wire.go:

var UserDependencySet = wire.NewSet(mySql.NewUserPersistenceRepository, usecase.NewUserUseCase, http.NewUserCommandHandler, http.NewUserQueryHandler)
