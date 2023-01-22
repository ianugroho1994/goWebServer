package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	orgHandler "hardtmann/smartlab/api/organization/handler"
	orgRepo "hardtmann/smartlab/api/organization/repository"

	userHandler "hardtmann/smartlab/api/user/handler"
	userRepo "hardtmann/smartlab/api/user/repository"
	"hardtmann/smartlab/database"

	"github.com/iancoleman/strcase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func configureMiddleware(e *echo.Echo) {
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		Skipper: func(c echo.Context) bool {
			// use skipper to prevent back and forth redirection using staticfs
			return strings.Contains(c.Request().RequestURI, "admin")
		},
	}))
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(100)))

	// jwtConfig := middleware.JWTConfig{
	// 	Claims:      &helpers.SmartLabJwtClaims{},
	// 	SigningKey:  []byte(viper.GetString("AES_PKEY_32BIT")),
	// 	ContextKey:  "user-jwt",
	// 	TokenLookup: "header:Authorization",
	// 	AuthScheme:  "Bearer",
	// 	Skipper: func(c echo.Context) bool {
	// 		return strings.Contains(c.Request().RequestURI, "login") ||
	// 			strings.Contains(c.Request().RequestURI, "callback") ||
	// 			strings.Contains(c.Request().RequestURI, "admin") ||
	// 			strings.Contains(c.Request().RequestURI, "webapp") ||
	// 			strings.Contains(c.Request().RequestURI, "favicon") || strings.Contains(c.Request().RequestURI, "alarm/files")
	// 	},
	// }

	// e.Use(middleware.JWTWithConfig(jwtConfig))
}

func configureErrorHandler(e *echo.Echo) {
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if castedObject, ok := err.(validator.ValidationErrors); ok {
			msgs := []string{}
			for _, err := range castedObject {
				errField := strcase.ToSnake(err.Field())
				errParam := strcase.ToSnake(err.Param())
				switch err.Tag() {
				case "required_with":
					msgs = append(msgs, fmt.Sprintf("%s is required with %s",
						errField, errParam))
				case "url":
					msgs = append(msgs, fmt.Sprintf("%s is not a valid url",
						errField))
				case "oneof":
					msgs = append(msgs, fmt.Sprintf("%s must be one of %s",
						errField, errParam))
				case "required":
					msgs = append(msgs, fmt.Sprintf("%s is required",
						errField))
				case "email":
					msgs = append(msgs, fmt.Sprintf("%s is not valid email",
						errField))
				case "gte":
					msgs = append(msgs, fmt.Sprintf("%s value must be greater than %s",
						errField, errParam))
				case "lte":
					msgs = append(msgs, fmt.Sprintf("%s value must be lower than %s",
						errField, errParam))
				}
			}
			report.Message = strings.Join(msgs, ", ")
		}

		c.Logger().Error(report)
		c.JSON(report.Code, report)
	}
}

func configureValidator(e *echo.Echo) {
	// e.Validator = &common.CustomValidator{
	// 	Validator: validator.New(),
	// }
}

func NewAPIServer() *echo.Echo {
	e := echo.New()
	timeout := time.Duration(viper.GetInt("context.timeout")) * time.Second

	//configure middleware and error handling
	configureErrorHandler(e)
	configureMiddleware(e)
	configureValidator(e)

	databaseConnection := database.GetDB()

	//init API and it's dependencies inside
	organizationAPI := orgHandler.NewOrganizationAPI(orgRepo.NewOrganizationRepository(databaseConnection))
	userAPI := userHandler.NewUserAPI(userRepo.NewUserRepository(databaseConnection), timeout)

	//populate api routes
	organizationAPI.PopulateRouteHandler(e)
	userAPI.PopulateRouteHandler(e)
	log.Println("finish init api server")
	return e

}
