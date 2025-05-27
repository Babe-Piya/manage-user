package controller

import (
	"go.uber.org/zap"
	"manage-user/services"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	CreateUser(c echo.Context) error
	GetListUser(c echo.Context) error
	GetUserByID(c echo.Context) error
	UpdateUserByID(c echo.Context) error
	DeleteUserByID(c echo.Context) error
	Login(c echo.Context) error
}
type userController struct {
	UserService services.UserService
	Log         *zap.Logger
}

func NewUserController(userService services.UserService, log *zap.Logger) UserController {
	return &userController{
		UserService: userService,
		Log:         log,
	}
}
