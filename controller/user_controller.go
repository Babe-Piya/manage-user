package controller

import (
	"manage-user/services"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	CreateUser(c echo.Context) error
}
type userController struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return &userController{
		UserService: userService,
	}
}
