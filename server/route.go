package server

import (
	"net/http"

	"manage-user/appconfig"
	"manage-user/common"
	"manage-user/controller"
	"manage-user/middlewares"
	"manage-user/repositories"
	"manage-user/services"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func routes(e *echo.Echo, db *mongo.Database, config *appconfig.AppConfig) {
	userRepo := repositories.NewUserRepository(db)

	userSrv := services.NewUserService(userRepo, config)

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
	userAPI.POST("/login", userCtrl.Login)
	userAPI.POST("/create", userCtrl.CreateUser)

	auth := middlewares.NewAuthorization(config.JwtSecret)
	userAPI.Use(auth.AuthorizationMiddleware)
	userAPI.POST("/update", userCtrl.UpdateUserByID)
	userAPI.GET("/list", userCtrl.GetListUser)
	userAPI.GET("/:id", userCtrl.GetUserByID)
	userAPI.DELETE("/delete/:id", userCtrl.DeleteUserByID)
}
