package config

import "context"

type Config struct {
	count int

	interval int
	cfgPath  string

	context    context.Context
	cancelFunc context.CancelFunc
}

type Option func(ctx context.Context, cfg *Config)

func New(opts ...Option) *Config {
	cfg := &Config{
		count: 1,
	}
	cfg.context, cfg.cancelFunc = context.WithCancel(context.Background())

	for _, opt := range opts {
		opt(cfg.context, cfg)
	}

	return cfg
}
