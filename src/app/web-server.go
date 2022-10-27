package app

import (
	"github.com/labstack/echo"
	http2 "goWebServer/internal/users/delivery/http"
)

type WebServer struct {
	EchoInstance       *echo.Echo
	UserCommandHandler *http2.UserCommandHandler
	UserQueryHandler   *http2.UserQueryHandler
}

func NewUserServer(e *echo.Echo, c *http2.UserCommandHandler, q *http2.UserQueryHandler) *WebServer {
	return &WebServer{
		UserCommandHandler: c,
		UserQueryHandler:   q,
		EchoInstance:       e,
	}
}
