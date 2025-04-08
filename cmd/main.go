package main

import (
	"context"
	"github.com/th1nksnow/thehousewolf/internal/api/bot"
	"github.com/th1nksnow/thehousewolf/internal/config"
	"github.com/th1nksnow/thehousewolf/internal/openai"
	"github.com/th1nksnow/thehousewolf/internal/redisdb"
	"log/slog"
	"os"
)

func main() {
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	ctx := context.Background()

	cfg := config.InitConfig()

	log.Info("init config success")
	log.Debug(cfg.Bot.Name)

	redisCfg := redisdb.NewConfig(cfg.Redis)
	db, err := redisdb.InitRedis(ctx, redisCfg)
	if err != nil {
		log.Error("failed to init redis client", err)
		os.Exit(1)
	}
	log.Info("init redis success")

	redisMsg := redisdb.NewRedis(db)

	openaiCfg := openai.NewConfig(cfg.OpenAI.Key)
	openaiService := openai.NewService(openaiCfg)

	botCfg := bot.NewConfig(cfg.Bot.Token, cfg.Bot.Name)
	botWrapper, err := bot.NewWrapper(botCfg, openaiService, redisMsg)
	if err != nil {
		log.Error("failed to create bot wrapper", err)
		os.Exit(1)
	}
	log.Info("init bot success")

	err = botWrapper.Start(ctx)
	if err != nil {
		log.Error("failed to start bot wrapper", err)
		os.Exit(1)
	}
	log.Info("bot run success")
	os.Exit(0)
}
