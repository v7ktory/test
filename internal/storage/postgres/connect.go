package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/v7ktory/test/internal/config"
)

type Postgres struct {
	Pool *sql.DB
}

func New(cfg *config.Config) *Postgres {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Database, cfg.Postgres.SSLmode)
	pool, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("unable to use data source name", err)
	}

	pool.SetConnMaxLifetime(0)
	pool.SetMaxIdleConns(3)
	pool.SetMaxOpenConns(3)

	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	Ping(ctx, pool)

	return &Postgres{Pool: pool}
}

func Ping(ctx context.Context, pool *sql.DB) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err := pool.PingContext(ctx); err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
}
