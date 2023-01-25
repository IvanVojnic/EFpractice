package repository

import (
	"EFpractic2/pkg/utils"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"

	"time"
)

type StorageConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DBName   string `json:"database"`
}

func NewPostgresDB(cfg StorageConfig, ctx context.Context, maxAttempts int) (pool *pgxpool.Pool, err error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.Username, cfg.Password, cfg.Port, cfg.DBName)
	utils.DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		pool, err = pgxpool.Connect(ctx, dsn)
		if err != nil {
			return err
		}
		return nil
	}, maxAttempts, 5*time.Second)
	if err != nil {
		log.Fatal("error to connect")
	}
	return pool, err
}
