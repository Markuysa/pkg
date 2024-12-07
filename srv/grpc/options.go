package grpc

import (
	"fmt"
	"time"
)

type regFuncs []func()

func (r regFuncs) apply() {
	for _, f := range r {
		f()
	}
}

type opts struct {
	config   *Config
	regFuncs regFuncs
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
	f []func()
}

func (fo *funcOption) apply(o *opts) {
	o.regFuncs = fo.f
}

func WithRegisterFuncs(f ...func()) option {
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
