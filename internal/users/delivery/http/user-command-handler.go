package http

import (
	"goWebServer/shared/domain/user"
	"goWebServer/shared/helpers"
	"goWebServer/shared/logger"
	"goWebServer/shared/properties"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type UserCommandHandler struct {
	userUseCase user.UserUseCase
	validate    *validator.Validate
}

func NewUserCommandHandler(echo *echo.Echo, userUC user.UserUseCase) *UserCommandHandler {
	handler := &UserCommandHandler{
		userUseCase: userUC,
		validate:    validator.New(),
	}

	userGroup := echo.Group("api/user")
	userGroup.PUT("/:id", handler.UpdateByID)
	return handler
}

func (handler *UserCommandHandler) UpdateByID(c echo.Context) (err error) {
	var Request UserUpdateRequest
	id := c.Param("id")

	err = helpers.BindRequest(c, UserUpdateRequest{})
	if err != nil {
		return err
	}

	err = helpers.ValidateStruct(c, Request)
	if err != nil {
		return err
	}

	startingDateParse, _ := time.Parse("YYYY-MM-DD", Request.StartingDate)
	userProfile := user.User{
		ID:           id,
		Name:         Request.Name,
		Address:      Request.Address,
		Email:        Request.Email,
		PhoneNumber:  Request.PhoneNumber,
		StartingDate: startingDateParse,
	}

	//check if citizen detail is available
	userProfile, err = handler.userUseCase.GetByID(c.Request().Context(), userProfile.ID)
	if err != nil {
		logger.Log.Fatal(err.Error())
		return c.JSON(http.StatusBadRequest, properties.ErrorResponse{Message: err.Error(), Data: err})

	}

	userProfile, err = handler.userUseCase.UpdateProfile(c.Request().Context(), &userProfile)
	if err != nil {
		logger.Log.Fatal(err.Error())
		return c.JSON(http.StatusUnauthorized, properties.ErrorResponse{Message: err.Error(), Data: err})
	}

	return c.JSON(http.StatusOK, userProfile)
}
