package grpc

import (
	"fmt"
	"time"

	"google.golang.org/grpc"
)

type registerer interface {
	RegisterServer(server *grpc.Server)
}

type regs []registerer

func (r regs) apply(server *grpc.Server) {
	for _, reg := range r {
		reg.RegisterServer(server)
	}
}

type opts struct {
	config *Config
	regs   regs
}

func (o opts) Validate() error {
	if o.config == nil {
		return fmt.Errorf("config is required")
	}

	return nil
}

type option interface {
	apply(*opts)
}

type funcOption struct {
	f []registerer
}

func (fo *funcOption) apply(o *opts) {
	o.regs = fo.f
}

func WithRegistes(f ...registerer) option {
	return &funcOption{f: f}
}

type cfgOption struct {
	c *Config
}

func (co *cfgOption) apply(o *opts) {
	o.config = co.c
}

func WithConfig(c *Config) option {
	return &cfgOption{c: c}
}

func defaultOptions() opts {
	return opts{
		config: &Config{
			Host:              "0.0.0.0:9000",
			MaxConnectionIdle: 5 * time.Minute,
			MaxConnectionAge:  5 * time.Minute,
			Timeout:           5 * time.Minute,
			Time:              5 * time.Minute,
		},
	}
}
