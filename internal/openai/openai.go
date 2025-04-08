package openai

import (
	"context"
	"github.com/cohesion-org/deepseek-go"
	"log/slog"
)

const (
	openAiQwen   = "deepseek/deepseek-r1-distill-qwen-32b:free"
	openAiR1     = "deepseek/deepseek-r1:free"
	openAiLlama  = "deepseek/deepseek-r1-distill-llama-70b:free"
	openAiGemini = "google/gemma-3-27b-it:free"
)

type OpenAiService struct {
	client *deepseek.Client
}

func NewOpenAiService(cfg *Config) *OpenAiService {
	err := cfg.Validate()
	if err != nil {
		slog.Error("openai: invalid config", err)
		return nil
	}

	baseURL := "https://openrouter.ai/api/v1/"
	client := deepseek.NewClient(cfg.Key, baseURL)

	return &OpenAiService{client: client}
}

func (s *OpenAiService) ChatCompetition(ctx context.Context, chat []string) (string, error) {
	messages := []deepseek.ChatCompletionMessage{
		{
			Role: deepseek.ChatMessageRoleSystem,
			Content: `
Тебя зовут Волк. Ты бот-помошник Telegram в беседе, состоящей из несекольких участников.
Никнеймы участников пишутся в начале сообщения.
Не повторяй сообщений участников, если они этого не попросят.
Отправляй ответ только на последний запрос, но не рассказывай об этом.`,
		},
	}

	for _, msg := range chat {
		messages = append(messages, deepseek.ChatCompletionMessage{
			Role:    deepseek.ChatMessageRoleUser,
			Content: msg,
		})
	}

	response, err := s.client.CreateChatCompletion(ctx, &deepseek.ChatCompletionRequest{
		Model:    openAiGemini,
		Messages: messages,
	})
	if err != nil {
		slog.Error("failed to response", err)
		return "", err
	}

	return response.Choices[0].Message.Content, err
}
