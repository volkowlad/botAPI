package openai

import (
	"errors"
	"log/slog"
)

type Config struct {
	Key string
}

func NewConfig(key string) *Config {
	return &Config{
		Key: key,
	}
}

func (c *Config) Validate() error {
	if c.Key == "" {
		slog.Error("openai: missing openai key")
		return errors.New("openai: missing openai key")
	}

	return nil
}
