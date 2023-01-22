package handler

import (
	"hardtmann/smartlab/api/organization/domain"

	"github.com/labstack/echo/v4"
)

type OrganizationAPI struct {
	commands *OrganizationAPICommands
	queries  *OrganizationAPIQueries
}

func NewOrganizationAPI(orgRepo domain.IOrganizationRepository) *OrganizationAPI {
	return &OrganizationAPI{
		commands: NewOrganizationAPICommand(orgRepo),
		queries:  NewOrganizationAPIQueries(orgRepo),
	}
}

func (o *OrganizationAPI) PopulateRouteHandler(e *echo.Echo) {
	route := e.Group("/org")
	o.commands.PopulateCommandAPI(route)
	o.queries.PopulateQueriesAPI(route)
}
