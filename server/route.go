package server

import (
	"manage-user/logger"
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
	"go.uber.org/zap"
)

func routes(e *echo.Echo, db *mongo.Database, config *appconfig.AppConfig, log *zap.Logger) {
	userRepo := repositories.NewUserRepository(db)

	userSrv := services.NewUserService(userRepo, config, log)

	userCtrl := controller.NewUserController(userSrv, log)

	// Custom Validator
	// Initialize the validator
	e.Validator = &common.CustomValidator{Validator: validator.New()}

	e.GET("/health", func(c echo.Context) error {
		response := map[string]string{
			"message": "service available",
		}
		return c.JSON(http.StatusOK, response)
	})

	e.Use(logger.ZapLogger(log))

	userAPI := e.Group("/user")
	userAPI.POST("/login", userCtrl.Login)
	userAPI.POST("/register", userCtrl.CreateUser)

	auth := middlewares.NewAuthorization(config.JwtSecret)
	userAPI.Use(auth.AuthorizationMiddleware)
	userAPI.POST("/update", userCtrl.UpdateUserByID)
	userAPI.GET("/list", userCtrl.GetListUser)
	userAPI.GET("/:id", userCtrl.GetUserByID)
	userAPI.DELETE("/delete/:id", userCtrl.DeleteUserByID)
}
