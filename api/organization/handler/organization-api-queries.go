package handler

import (
	"hardtmann/smartlab/api/organization/domain"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrganizationAPIQueries struct {
	repo domain.IOrganizationRepository
}

func NewOrganizationAPIQueries(orgRepo domain.IOrganizationRepository) *OrganizationAPIQueries {
	return &OrganizationAPIQueries{
		repo: orgRepo,
	}
}

func (o *OrganizationAPIQueries) PopulateQueriesAPI(g *echo.Group) {
	g.GET("/:id", o.GetByID)
	g.GET("/all", o.GetAll)
}

func (o *OrganizationAPIQueries) GetByID(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	o.repo.GetByID(id)
	return nil
}

func (o *OrganizationAPIQueries) GetAll(c echo.Context) error {
	_, err := o.repo.GetAll()
	if err != nil {
		return err
	}
	return nil
}
