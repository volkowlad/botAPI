package openai

import (
	"context"
)

type OpenAi interface {
	ChatCompetition(ctx context.Context, chat []string) (string, error)
}

type Service struct {
	OpenAi
}

func NewService(cfg *Config) *Service {
	return &Service{
		OpenAi: NewOpenAiService(cfg),
	}
}
