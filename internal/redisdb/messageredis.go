package redisdb

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"time"
)

type MessageRedis struct {
	db *redis.Client
}

func NewMessageRedis(db *redis.Client) *MessageRedis {
	return &MessageRedis{db: db}
}

func (m *MessageRedis) AddMessage(ctx context.Context, key, message string) error {
	err := m.db.RPush(ctx, key, message).Err()
	if err != nil {
		return err
	}

	err = m.db.Expire(ctx, key, time.Hour*48).Err()
	if err != nil {
		return err
	}

	return nil
}

func (m *MessageRedis) GetChat(ctx context.Context, key string) ([]string, error) {
	chat, err := m.db.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	return chat, nil
}

func (m *MessageRedis) Delete(ctx context.Context, key string) error {
	exists, err := m.db.Exists(ctx, key).Result()
	if err != nil || exists == 0 {
		slog.Error(fmt.Sprintf("key does not exist: %v", err))
		return errors.New("key does not exist")
	}

	err = m.db.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	return nil
}
