package redisdb

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log/slog"
)

func InitRedis(ctx context.Context, cfg *Config) (*redis.Client, error) {
	err := cfg.Validate()
	if err != nil {
		slog.Error("Init Redis Config Error", err)
		return nil, err
	}

	db := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if err := db.Ping(ctx).Err(); err != nil {
		slog.Error("failed to connect to redis server")
		return nil, err
	}

	return db, nil
}
