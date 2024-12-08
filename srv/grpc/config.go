package grpc

import "time"

type Config struct {
	Host              string        `default:"0.0.0.0:9000" yaml:"host"`
	MaxConnectionIdle time.Duration `default:"5m" yaml:"max_connection_idle"`
	MaxConnectionAge  time.Duration `default:"5m" yaml:"max_connection_age"`
	Timeout           time.Duration `default:"5m" yaml:"timeout"`
	Time              time.Duration `default:"5m" yaml:"time"`
	MaxRecvMsgSize    int           `default:"10485760" yaml:"max_recv_msg_size"`
	MaxSendMsgSize    int           `default:"10485760" yaml:"max_send_msg_size"`
}
