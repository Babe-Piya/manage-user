package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"manage-user/appconfig"
	"manage-user/database"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func Start(config *appconfig.AppConfig) (*echo.Echo, *mongo.Database) {
	// mongodb://user:password@host:port
	db, err := database.NewConnection(&config.MongoDB)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	routes(e, db, config)

	go func() {
		endPoint := fmt.Sprintf(":%s", config.ServerPort)
		if err := e.Start(endPoint); err != nil && !errors.Is(err, http.ErrServerClosed) {
			e.Logger.Error(err.Error())
			e.Logger.Fatal("shutting down the server")
		}
	}()

	return e, db
}

func Shutdown(e *echo.Echo, db *mongo.Database) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	defer func() {
		if err := db.Client().Disconnect(ctx); err != nil {
			e.Logger.Fatal(err)
		}
	}()
}
