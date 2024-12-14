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
		Port     int           `default:"5432"`
		User     string        `default:"postgres"`
		Password string        `default:"postgres"`
		Database string        `default:"postgres"`
		SSLMode  string        `default:"disable"`
		Timeout  time.Duration `default:"5s"`
		Extra    Extra
		Migrate  *MigrateCfg
	}
	Extra struct {
		MaxOpenConnections int32 `default:"10"`
		MinOpenConnections int32 `default:"1"`
	}
)

func New(
	c PgxPoolCfg,
	opts ...option,
) (*pgxpool.Pool, error) {
	opt := &options{}
	for _, o := range opts {
		o.apply(opt)
	}

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

	if opt.migrate != nil {
		err = applyMigrations(opt.migrate)
		if err != nil {
			return nil, fmt.Errorf("failed to launch migrate: %v", err)
		}
	}

	return pool, err
}
