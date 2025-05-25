package server

import (
	"net/http"

	"manage-user/common"
	"manage-user/controller"
	"manage-user/repositories"
	"manage-user/services"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func routes(e *echo.Echo, db *mongo.Database) {
	userRepo := repositories.NewUserRepository(db)

	userSrv := services.NewUserService(userRepo)

	userCtrl := controller.NewUserController(userSrv)

	// Custom Validator
	// Initialize the validator
	e.Validator = &common.CustomValidator{Validator: validator.New()}

	e.GET("/health", func(c echo.Context) error {
		response := map[string]string{
			"message": "service available",
		}
		return c.JSON(http.StatusOK, response)
	})

	userAPI := e.Group("/user")
	userAPI.POST("/create", userCtrl.CreateUser)
	userAPI.GET("/list", userCtrl.GetListUser)
	userAPI.GET("/:id", userCtrl.GetUserByID)
}
