package grpc

import "time"

type Config struct {
	Host              string        `default:"0.0.0.0:9000"`
	MaxConnectionIdle time.Duration `default:"5m"`
	MaxConnectionAge  time.Duration `default:"5m"`
	Timeout           time.Duration `default:"5m"`
	Time              time.Duration `default:"5m"`
	MaxRecvMsgSize    int
	MaxSendMsgSize    int
}
