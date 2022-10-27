package helpers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"goWebServer/shared/properties"
	"net/http"
)

func BindRequest(c echo.Context, request interface{}) (err error) {
	err = c.Bind(&request)
	if err != nil {
		//logger.Log.Fatal(err.Error())
		return c.JSON(http.StatusBadRequest, properties.ErrorResponse{Message: err.Error(), Data: err})
	}
	return nil
}

func ValidateStruct(c echo.Context, request interface{}) (err error) {
	err = validator.New().Struct(request)
	if err != nil {
		//logger.Log.Fatal(err.Error())
		return c.JSON(http.StatusBadRequest, properties.ErrorResponse{Message: err.Error(), Data: err})
	}
	return nil
}
