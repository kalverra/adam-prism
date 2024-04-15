package main

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	LogLevel  string  `env:"LOG_LEVEL" envDefault:"debug"`
	Latitude  float64 `env:"LATITUDE"`
	Longitude float64 `env:"LONGITUDE"`
	NotionKey string  `env:"NOTION_KEY"`
}

func ReadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		if err.Error() != "open .env: no such file or directory" {
			return nil, err
		}
		log.Warn().Msg(".env file not found, if running in production this probably is expected")
	} else {
		log.Info().Msg("Loaded .env file")
	}
	config := &Config{}
	err = env.Parse(config)
	return config, err
}
