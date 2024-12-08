package redis

import "time"

type Config struct {
	Host     string        `validate:"required" envconfig:"REDIS_HOST" default:"localhost"`
	Port     int           `validate:"required" envconfig:"REDIS_PORT" default:"6379"`
	Password string        `validate:"required" envconfig:"REDIS_PASSWORD"`
	DB       int           `envconfig:"REDIS_DB" default:"0"`
	Timeout  time.Duration `validate:"required" envconfig:"REDIS_TIMEOUT" default:"5s"`
}
