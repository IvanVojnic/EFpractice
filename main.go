package main

import (
	"EFpractic2/pkg/config"
	"EFpractic2/pkg/handler"
	"EFpractic2/pkg/repository"
	"EFpractic2/pkg/service"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	cfg := config.GetConfig()
	db, err := repository.NewPostgresDB(cfg.Storage)
	if err != nil {
		fmt.Sprintf("error get db: %s", err)
	}
	defer repository.ClosePool(db)

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	handlers.InitRoutes(e)
}

/*
func DBConnection(Cfg *Config) (Repository, error) {
	switch Cfg.CurrentDB {
	case "postgres":
		pool, err := pgxpool.New(context.Background(), Cfg.PostgresUrl)
		if err != nil {
			return nil, fmt.Errorf("invalid configuration data: %v", err)
		}
		if err = pool.Ping(context.Background()); err != nil {
			return nil, fmt.Errorf("database not responding: %v", err)
		}
		return &PRepository{Pool: pool}, nil
	case "mongo":
	}
	return nil, nil
}

func ClosePool(Cfg *Config, r interface{}) {
	switch Cfg.CurrentDB {
	case "postgres":
		pr := r.(PRepository)
		if pr.Pool != nil {
			pr.Pool.Close()
		}
	case "mongo":
	}
}*/
