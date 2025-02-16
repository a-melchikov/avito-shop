package db

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"

	"github.com/a-melchikov/avito-shop/internal/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPostgresDB(cfg *config.Config) *pgxpool.Pool {
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		logrus.Fatalf("Ошибка подключения к БД: %v", err)
	}

	if err := pool.Ping(ctx); err != nil {
		logrus.Fatalf("Невозможно подключиться к БД: %v", err)
	}

	return pool
}
