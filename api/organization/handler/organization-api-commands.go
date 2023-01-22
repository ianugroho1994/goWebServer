package handler

import (
	"hardtmann/smartlab/api/organization/domain"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrganizationAPICommands struct {
	repo domain.IOrganizationRepository
}

func NewOrganizationAPICommand(orgRepo domain.IOrganizationRepository) *OrganizationAPICommands {
	return &OrganizationAPICommands{
		repo: orgRepo,
	}
}

func (o *OrganizationAPICommands) PopulateCommandAPI(g *echo.Group) {
	g.POST("", o.Add)
	g.DELETE("/:id", o.Delete)
	g.PUT("", o.Update)
}

func (o *OrganizationAPICommands) Add(c echo.Context) error {
	o.repo.Create(&domain.Organization{})
	return nil
}

func (o *OrganizationAPICommands) Delete(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	o.repo.Delete(id)
	return nil
}

func (o *OrganizationAPICommands) Update(c echo.Context) error {
	o.repo.Update(&domain.Organization{})
	return nil
}
