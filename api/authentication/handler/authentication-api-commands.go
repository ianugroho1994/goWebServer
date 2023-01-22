package handler

import (
	"errors"
	"hardtmann/smartlab/api/authentication/domain"
	userDomain "hardtmann/smartlab/api/user/domain"

	"github.com/aeeem/utilities"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type AuthenticationAPICommands struct {
	useCase domain.IAuthenticationUseCase
}

func NewAuthenticationAPICommand(userUseCase domain.IAuthenticationUseCase) *AuthenticationAPICommands {
	return &AuthenticationAPICommands{
		useCase: userUseCase,
	}
}

func (o *AuthenticationAPICommands) PopulateCommandAPI(g *echo.Group) {
	g.PUT("/login", o.Login)
}

func (o *AuthenticationAPICommands) Login(c echo.Context) error {
	var req LoginRequest
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	deviceID := c.Request().Header.Get("X-Device-Id")

	//validating request
	if deviceID == "" {
		err = errors.New("Given param is not valid")
		return err
	}

	//validating request
	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		return err
	}

	User := userDomain.User{
		Email:    req.Email,
		Password: req.Password,
	}

	ctx := c.Request().Context()
	auth, err := o.useCase.Login(ctx, &User, deviceID)
	if err != nil {
		return err
	}

	auth.User.Password = ""

	//c.Response().Header().Set("X-CSRF-Token", csrf)
	return c.JSON(utilities.StandardResponse(auth, "ok"))
}
