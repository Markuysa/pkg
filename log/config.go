package log

type Config struct {
	Level string `envconfig:"LOG_LEVEL" default:"info"`
}
