package http

import (
	"goWebServer/shared/domain/admin"
	"goWebServer/shared/domain/user"
	"goWebServer/shared/helpers"
	"goWebServer/shared/logger"
	"goWebServer/shared/properties"
	
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type AdminCommandHandler struct{
	adminUseCase admin.AdminUseCase
	validate *validator.Validate
}

func NewAdminCommandHandler(echo *echo.Echo, adminUC admin.AdminUseCase) *AdminCommandHandler{
	handler := &AdminCommandHandler{
		adminUseCase: adminUC,
		validate: validator.New(),
	}

	adminGroup := echo.Group("/admin")
	//todo: add authorization middleware

	
	adminGroup.PUT("/:id",handler.UpdateAdminProfile).Name = "update admin profile"
	adminGroup.POST("/register",handler.Register).Name = "register new admin"
	adminGroup.POST("/login",handler.Login).Name = "admin login"
	adminGroup.POST("/refresh_token",handler.RefreshToken).Name="refresh token"
	adminGroup.DELETE("/:id",handler.Delete).Name = "delete other admin by id"
	return handler
}

func (handler AdminCommandHandler) UpdateAdminProfile(c echo.Context)(err error){
	var Request AdminUpdateRequest
	idString := c.Param("id")

	err = helpers.BindRequest(c, AdminUpdateRequest{})
	if err != nil {
		return err
	}

	err = helpers.ValidateStruct(c, Request)
	if err != nil {
		return err
	}

	id,err := strconv.ParseInt(idString, 10, 64)
	adminProfile, err := handler.adminUseCase.GetByID(c.Request().Context(), id )
	if err != nil {
		logger.Log.Fatal(err.Error())
		return c.JSON(http.StatusBadRequest, properties.ErrorResponse{Message: err.Error(), Data: err})
	}

	err = handler.adminUseCase.UpdateProfile(c.Request().Context(), &adminProfile)
	if err != nil {
		logger.Log.Fatal(err.Error())
		return c.JSON(http.StatusUnauthorized, properties.ErrorResponse{Message: err.Error(), Data: err})
	}

	return c.JSON(http.StatusOK, adminProfile)
}

func (handler AdminCommandHandler) Register(c echo.Context)(err error){
	var Request AdminRegisterRequest

	err = helpers.BindRequest(c, AdminRegisterRequest{})
	if err != nil {
		return err
	}

	err = helpers.ValidateStruct(c, Request)
	if err != nil {
		return err
	}

	userData := user.User{
		Email: Request.Email,
		Password: Request.Password,
		Name: Request.Name,
		Address: Request.Address,
		PhoneNumber: Request.Address,
	}

	adminData := admin.Admin{
		User: &userData,
	}

	err, statusCode := handler.adminUseCase.RegisterAdmin(c.Request().Context(), &adminData)
	if err != nil{
		return c.JSON(int(statusCode), properties.ErrorResponse{Message: err.Error(), Data: err})
	}

	return c.JSON(http.StatusOK, adminData)
}

func (impl AdminCommandHandler) Delete(c echo.Context)(err error){
	
	return nil
}

func (impl AdminCommandHandler) RefreshToken(c echo.Context)(err error){
	
	return nil
}

func (impl AdminCommandHandler) Login(c echo.Context)(err error){
	
	return nil
}

