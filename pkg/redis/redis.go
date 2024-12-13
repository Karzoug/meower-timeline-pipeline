package redis

import (
	"context"

	"gopkg.in/redis.v5"
)

func New(cfg Config) (*redis.Client, func(context.Context) error, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.Addr,
		DB:   cfg.Database,
	})

	if cmd := rdb.Ping(); cmd.Err() != nil {
		return nil, nil, cmd.Err()
	}

	return rdb, func(ctx context.Context) error { return rdb.Close() }, nil
}
