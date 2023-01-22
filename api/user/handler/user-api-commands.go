package handler

import (
	responseError "hardtmann/smartlab/api/response-error"
	"hardtmann/smartlab/api/user/domain"
	"hardtmann/smartlab/helpers"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type UserAPICommands struct {
	useCase domain.IUserUseCase
	repo    domain.IUserRepository
}

func NewUserAPICommand(userUseCase domain.IUserUseCase, userRepo domain.IUserRepository) *UserAPICommands {
	return &UserAPICommands{
		useCase: userUseCase,
		repo:    userRepo,
	}
}

func (o *UserAPICommands) PopulateCommandAPI(g *echo.Group) {
	g.PUT("/change_password", o.ChangePassword)
	g.POST("/create", o.CreateNewUser)
	g.DELETE("/:id", o.DeleteUser)
}

func (o *UserAPICommands) ChangePassword(c echo.Context) error {
	var UserID int64
	var Request ChangePasswordRequest

	UserID = c.(helpers.SmartLabJwtClaims).UserID
	c.Bind(&Request)

	validate := validator.New()
	err := validate.Struct(Request)
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = o.useCase.ChangePassword(c.Request().Context(), UserID, Request.Password)
	if err != nil {
		return err
	}

	return nil
}

// CreateNewUser should move this to auth
func (o *UserAPICommands) CreateNewUser(c echo.Context) (err error) {
	var Request CreateUserRequest
	err = c.Bind(&Request)
	if err != nil {
		log.Fatalln(err.Error())
		return c.JSON(http.StatusBadRequest, responseError.ResponseError{Message: err.Error(), Data: err})
	}

	validate := validator.New()
	err = validate.Struct(Request)
	if err != nil {
		log.Fatalln(err.Error())
		return c.JSON(http.StatusBadRequest, responseError.ResponseError{Message: err.Error(), Data: err})
	}

	user := domain.User{
		Username: Request.Name,
		Email:    Request.Email,
		Password: Request.Password,
	}

	err = o.repo.Create(c.Request().Context(), &user)
	if err != nil {
		log.Fatalln(err.Error())
		return c.JSON(http.StatusBadRequest, responseError.ResponseError{Message: err.Error(), Data: err})
	}

	return c.JSON(http.StatusOK, user)
}

// DeleteUser should move this to admin
func (o *UserAPICommands) DeleteUser(c echo.Context) (err error) {
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Fatalln(err.Error())
		return c.JSON(http.StatusInternalServerError, responseError.ResponseError{Message: err.Error(), Data: err})
	}

	err = o.repo.Delete(c.Request().Context(), int64(userId))

	if err != nil {
		log.Fatalln(err.Error())
		return c.JSON(http.StatusInternalServerError, responseError.ResponseError{Message: err.Error(), Data: err})
	}

	return c.JSON(http.StatusOK, "ok")
}
