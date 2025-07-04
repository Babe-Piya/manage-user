package controller

import (
	"net/http"

	"manage-user/appconstants"
	"manage-user/services"

	"github.com/labstack/echo/v4"
)

func (ctrl *userController) UpdateUserByID(c echo.Context) error {
	var req services.UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		ctrl.Log.Error(err.Error())
		errResp := appconstants.NewErrorResponse(err)

		return c.JSON(http.StatusBadRequest, errResp)
	}

	if err := c.Validate(&req); err != nil {
		ctrl.Log.Error(err.Error())
		errResp := appconstants.NewErrorResponse(err)

		return c.JSON(http.StatusBadRequest, errResp)
	}

	ctx := c.Request().Context()
	resp, err := ctrl.UserService.UpdateUserByID(ctx, req)
	if err != nil {
		errResp := appconstants.NewErrorResponse(err)

		return c.JSON(http.StatusInternalServerError, errResp)
	}

	return c.JSON(http.StatusOK, resp)
}
