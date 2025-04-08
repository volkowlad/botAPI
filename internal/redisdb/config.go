package redisdb

import (
	"botAPI/internal/config"
	"errors"
	"log/slog"
)

type Config struct {
	Addr     string
	Password string
	DB       int
}

func NewConfig(c config.RedisConfig) *Config {
	return &Config{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       c.DB,
	}
}

func (c *Config) Validate() error {
	if c.Addr == "" {
		slog.Error("redis: missing redis host")
		return errors.New("redis: redis host")
	}

	if c.Password == "" {
		slog.Error("redis: missing redis password")
		return errors.New("redis: redis password")
	}

	if c.DB < 0 {
		slog.Error("redis: invalid redis db")
		return errors.New("redis: invalid redis db")
	}

	return nil
}
