package controller

import (
	"manage-user/services"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	CreateUser(c echo.Context) error
	GetListUser(c echo.Context) error
	GetUserByID(c echo.Context) error
	UpdateUserByID(c echo.Context) error // TODO: ไม่ให้อัพเดตซ้ำกับ email คนอื่น
	DeleteUserByID(c echo.Context) error // TODO: อ่าน token ไม่ให้ user ลบตัวเอง
	Login(c echo.Context) error
}
type userController struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return &userController{
		UserService: userService,
	}
}
