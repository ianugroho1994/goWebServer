package authorization

import (
	"goWebServer/shared/domain/user"
	"goWebServer/shared/properties"
	"strings"
	"net/http"

	"github.com/labstack/echo"
)

type AuthorizationMiddleware struct {
	User user.User
}

// CORS will handle the CORS middleware
func (m *AuthorizationMiddleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		c.Response().Header().Set("Access-Control-Allow-Methods", "*")
		c.Response().Header().Set("Access-Control-Allow-Headers", "*")

		return next(c)
	}
}

func InitMiddleware() *AuthorizationMiddleware {
	return &AuthorizationMiddleware{}
}

func (m *AuthorizationMiddleware) Authorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenBearer := strings.Split(c.Request().Header.Get("Authorization"), "Bearer ")
		if len(tokenBearer) < 2 {
			return c.JSON(http.StatusInternalServerError, properties.ErrorResponse{Message: "Bearer token empty or less then expected", Data: ""})
		}

		
		return nil
	}
}