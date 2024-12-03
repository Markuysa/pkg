package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	PgxPoolCfg struct {
		Host     string
		Port     int
		User     string
		Password string
		Database string
		SSLMode  string
		Timeout  time.Duration
		Extra    Extra
	}
	Extra struct {
		MaxOpenConnections int32
		MinOpenConnections int32
	}
)

func New(c PgxPoolCfg) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.User, c.Password, c.Host, c.Port, c.Database, c.SSLMode)

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse cfg: %v", err)
	}

	poolConfig.MaxConns = c.Extra.MaxOpenConnections
	poolConfig.MinConns = c.Extra.MinOpenConnections

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create pool: %v", err)
	}

	if err = pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ping: %v", err)
	}

	return pool, err
}
