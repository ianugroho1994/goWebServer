package handler

import (
	responseError "hardtmann/smartlab/api/response-error"
	"hardtmann/smartlab/api/user/domain"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserAPIQueries struct {
	repo domain.IUserRepository
}

func NewUserAPIQueries(orgRepo domain.IUserRepository) *UserAPIQueries {
	return &UserAPIQueries{
		repo: orgRepo,
	}
}

func (o *UserAPIQueries) PopulateQueriesAPI(g *echo.Group) {
	g.GET("", o.GetUsers)
	g.GET("/:id", o.GetUserByID)
}

func (o *UserAPIQueries) GetUsers(c echo.Context) (err error) {
	res, err := o.repo.GetAll(c.Request().Context())
	if err != nil {
		log.Fatalln(err.Error())
		return c.JSON(http.StatusInternalServerError, responseError.ResponseError{Message: err.Error(), Data: err})
	}
	return c.JSON(http.StatusOK, res)
}

func (o *UserAPIQueries) GetUserByID(c echo.Context) (err error) {
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Fatalln(err.Error())
		return c.JSON(http.StatusInternalServerError, responseError.ResponseError{Message: err.Error(), Data: err})
	}

	user, err := o.repo.GetByID(c.Request().Context(), userId)

	if err != nil {
		log.Fatalln(err.Error())
		return c.JSON(http.StatusInternalServerError, responseError.ResponseError{Message: err.Error(), Data: err})
	}

	return c.JSON(http.StatusOK, user)
}
