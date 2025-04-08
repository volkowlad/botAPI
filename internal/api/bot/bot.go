package bot

import (
	"botAPI/internal/openai"
	"botAPI/internal/redisdb"
	"context"
	"gopkg.in/telebot.v4"
	"log/slog"
)

const (
	codeCall   = "Аууу"
	deleteCall = "Луна зашла"
)

type Wrapper struct {
	bot       *telebot.Bot
	config    *Config
	openaiSrv *openai.Service
	redisMsg  redisdb.Message
}

func NewWrapper(config *Config, openaiSrv *openai.Service, redisMsg redisdb.Message) (*Wrapper, error) {
	err := config.Validate()
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	settings := telebot.Settings{
		Token: config.Token,
	}

	bot, err := telebot.NewBot(settings)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	w := &Wrapper{
		bot:       bot,
		config:    config,
		openaiSrv: openaiSrv,
		redisMsg:  redisMsg,
	}

	//bot.Use(middleware.Logger())
	go w.prepare()

	return w, nil
}

func (w *Wrapper) Start(_ context.Context) error {
	w.bot.Start()

	return nil
}

func (w *Wrapper) prepare() {
	w.bot.Handle(codeCall, w.startHandler)
	w.bot.Handle(telebot.OnText, w.textHandler)
	w.bot.Handle(deleteCall, w.deleteHandler)
}
