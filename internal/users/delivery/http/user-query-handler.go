package http

import (
	"goWebServer/shared/domain/user"
	"goWebServer/shared/logger"
	"goWebServer/shared/properties"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type UserQueryHandler struct {
	userUseCase user.UserUseCase
	validate    *validator.Validate
}

func NewUserQueryHandler(echo *echo.Echo, userUC user.UserUseCase) *UserQueryHandler {
	handler := &UserQueryHandler{
		userUseCase: userUC,
		validate:    validator.New(),
	}

	userGroup := echo.Group("api/user")
	userGroup.GET("", handler.GetUsers)
	userGroup.GET("/:id", handler.GetByID)
	return handler
}

func (handler *UserQueryHandler) GetUsers(c echo.Context) (err error) {
	res, _, err := handler.userUseCase.Fetch(c.Request().Context())
	if err != nil {
		logger.Log.Fatal(err.Error())
		return c.JSON(http.StatusInternalServerError, properties.ErrorResponse{Message: err.Error(), Data: err})
	}
	return c.JSON(http.StatusOK, res)

}

func (handler *UserQueryHandler) GetByID(c echo.Context) (err error) {
	citizenId := c.Param("id")
	citizen, err := handler.userUseCase.GetByID(c.Request().Context(), citizenId)

	if err != nil {
		logger.Log.Fatal(err.Error())
		return c.JSON(http.StatusInternalServerError, properties.ErrorResponse{Message: err.Error(), Data: err})
	}

	return c.JSON(http.StatusOK, citizen)
}
