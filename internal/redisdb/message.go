package redisdb

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type Message interface {
	AddMessage(ctx context.Context, key, message string) error
	GetChat(ctx context.Context, key string) ([]string, error)
	Delete(ctx context.Context, key string) error
}

type Repos struct {
	Message
}

func NewRedis(client *redis.Client) *Repos {
	return &Repos{
		Message: NewMessageRedis(client),
	}
}
