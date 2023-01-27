package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StorageConfig struct {
	Postgre_url string `json:"pUrl"`
}

func NewPostgresDB() (pool *pgxpool.Pool, err error) {
	//pool, err = pgxpool.New(context.Background(), "postgres://postgres:postgres@host.docker.internal:5432/postgres?sslmode=disable")
	pool, err = pgxpool.New(context.Background(), "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
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
