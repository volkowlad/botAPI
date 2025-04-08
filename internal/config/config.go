package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log/slog"
	"os"
)

type Config struct {
	Bot    BotConfig `yaml:"bot"`
	OpenAI OpenAIConfig
	Redis  RedisConfig `yaml:"redis"`
}

type BotConfig struct {
	Token string
	Name  string `yaml:"name"`
}

type OpenAIConfig struct {
	Key string
}

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string
	DB       int `yaml:"db"`
}

func initViper() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}

func InitConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		slog.Error("Error loading config file", err)
		return nil
	}

	err = initViper()
	if err != nil {
		slog.Error("Error initializing config", err)
		return nil
	}

	return &Config{
		Bot: BotConfig{
			Token: os.Getenv("TOKEN"),
			Name:  viper.GetString("bot.name"),
		},

		OpenAI: OpenAIConfig{
			Key: os.Getenv("API_KEY"),
		},

		Redis: RedisConfig{
			Addr:     viper.GetString("redis.addr"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       viper.GetInt("redis.db"),
		},
	}
}
