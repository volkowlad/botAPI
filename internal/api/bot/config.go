package bot

import (
	"errors"
	"log/slog"
)

type Config struct {
	Token string
	Name  string
}

func NewConfig(token, name string) *Config {
	return &Config{
		Token: token,
		Name:  name,
	}
}

func (c *Config) Validate() error {
	if c.Token == "" {
		slog.Error("bot token is empty")
		return errors.New("token is required")
	}

	if c.Name == "" {
		slog.Error("bot name is empty")
		return errors.New("bot name is required")
	}

	return nil
}
