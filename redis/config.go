package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func New(
	cfg Config,
) (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	options := &redis.Options{
		Addr:     addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	}

	client := redis.NewClient(options)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to ping redis: %v", err)
	}

	return client, nil
}
