package config

import "os"

type Config struct {
	Port          string
	TelegramToken string
}

func NewConfig() *Config {
	return &Config{
		Port:          os.Getenv("PORT"),
		TelegramToken: os.Getenv("TOKEN_API"),
	}
}
