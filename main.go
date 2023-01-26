package main

import (
	"EFpractic2/pkg/config"
	"EFpractic2/pkg/handler"
	"EFpractic2/pkg/repository"
	"EFpractic2/pkg/service"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {

	e := echo.New()
	logger := log.New()
	logger.Out = os.Stdout
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			log.WithFields(log.Fields{
				"URI":    values.URI,
				"status": values.Status,
			}).Info("request")
			return nil
		},
	}))

	cfg := config.GetConfig()
	fmt.Sprintf("CFG IS - %s", cfg)
	db, err := repository.NewPostgresDB(cfg.Storage)
	if err != nil {
		fmt.Sprintf("error get db: %s", err)
	}
	//defer repository.ClosePool(db)

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	handlers.InitRoutes(e)
}
