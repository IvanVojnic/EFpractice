package main

import (
	"EFpractic2/pkg/config"
	"EFpractic2/pkg/handler"
	"EFpractic2/pkg/repository"
	"EFpractic2/pkg/service"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	cfg := config.GetConfig()
	db, err := repository.NewPostgresDB(cfg.Storage, context.Background(), 5)
	if err != nil {
		fmt.Sprintf("error get db: %s", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	handlers.InitRoutes(e)
}
