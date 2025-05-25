package controller

import (
	"log"
	"net/http"

	"manage-user/appconstants"

	"github.com/labstack/echo/v4"
)

func (ctrl *userController) DeleteUserByID(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	resp, err := ctrl.UserService.DeleteUserByID(ctx, id)
	if err != nil {
		log.Println(err)
		errResp := appconstants.NewErrorResponse(err)

		return c.JSON(http.StatusInternalServerError, errResp)
	}

	return c.JSON(http.StatusOK, resp)
}
