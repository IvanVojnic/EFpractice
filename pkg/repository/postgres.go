package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StorageConfig struct {
	Postgre_url string `json:"pUrl"`
}

func NewPostgresDB(cfg StorageConfig) (pool *pgxpool.Pool, err error) {
	pool, err = pgxpool.New(context.Background(), cfg.Postgre_url)
	if err != nil {
		return nil, fmt.Errorf("invalid configuration data: %v", err)
	}
	if err = pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("database not responding: %v", err)
	}
	return pool, err
}

func ClosePool(myPool *pgxpool.Pool) {
	if myPool != nil {
		myPool.Close()
	}
}
