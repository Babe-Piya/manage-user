package controller

import (
	"log"
	"net/http"

	"manage-user/appconstants"
	"manage-user/services"

	"github.com/labstack/echo/v4"
)

func (ctrl *userController) CreateUser(c echo.Context) error {
	var req services.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		log.Println(err)
		errResp := appconstants.NewErrorResponse(err)

		return c.JSON(http.StatusBadRequest, errResp)
	}

	if err := c.Validate(&req); err != nil {
		log.Println(err)
		errResp := appconstants.NewErrorResponse(err)

		return c.JSON(http.StatusBadRequest, errResp)
	}

	ctx := c.Request().Context()
	resp, err := ctrl.UserService.CreateUser(ctx, req)
	if err != nil {
		log.Println(err)
		errResp := appconstants.NewErrorResponse(err)

		return c.JSON(http.StatusInternalServerError, errResp)
	}

	return c.JSON(http.StatusCreated, resp)
}
