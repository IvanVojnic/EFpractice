package main

import (
	"EFpractic2/pkg/config"
	"EFpractic2/pkg/handler"
	"EFpractic2/pkg/repository"
	"EFpractic2/pkg/service"
	"context"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/labstack/echo/v4"
)

type (
	Config struct {
		CurrentDB   string `env:"CURRENT_DB,notEmpty" envDefault:"postgres"`
		PostgresUrl string `env:"POSTGRES_DB_URL,notEmpty"`
		MongoURL    string `env:"MONGO_DB_URL,notEmpty"`
		JwtKey      string `env:"JWT_KEY,notEmpty"`
	}
)

func NewConfig() (*Config, error) {
	Cfg := &Config{}
	if err := env.Parse(Cfg); err != nil {
		return nil, fmt.Errorf("config - NewConfig: %v", err)
	}

	//Cfg.CurrentDB = "postgres"
	//Cfg.PostgresUrl = "postgres://postgres:postgres@host.docker.internal:5432/entity?sslmode=disable"
	//Cfg.MongoURL = "_"
	//Cfg.JwtKey = "874967EC3EA3490F8F2EF6478B72A756"
	return Cfg, nil
}

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
}
