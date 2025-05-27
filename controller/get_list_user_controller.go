package controller

import (
	"net/http"

	"manage-user/appconstants"

	"github.com/labstack/echo/v4"
)

func (ctrl *userController) GetListUser(c echo.Context) error {
	ctx := c.Request().Context()
	resp, err := ctrl.UserService.GetListUser(ctx)
	if err != nil {
		errResp := appconstants.NewErrorResponse(err)

		return c.JSON(http.StatusInternalServerError, errResp)
	}

	return c.JSON(http.StatusOK, resp)
}
