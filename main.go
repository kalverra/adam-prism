package main

import (
	"os"
	"time"

	"github.com/caarlos0/env/v10"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var config *Config

type Config struct {
	LogLevel  string  `env:"LOG_LEVEL" envDefault:"debug"`
	Latitude  float64 `env:"LATITUDE"`
	Longitude float64 `env:"LONGITUDE"`
	NotionKey string  `env:"NOTION_KEY"`
}

// read config and setup logger
func init() {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "15:04:05.00", // hh:mm:ss.ss format
	})

	err := godotenv.Load()
	if err != nil {
		if err.Error() != "open .env: no such file or directory" {
			log.Warn().Err(err).Msg("Failed to load .env file")
		} else {
			log.Warn().Msg(".env file not found, if running in production this probably is expected")
		}
	} else {
		log.Info().Msg("Loaded .env file")
	}
	config = &Config{}
	if err := env.Parse(config); err != nil {
		log.Fatal().Err(err).Msg("Failed to parse environment variables")
	}

	logLvl, err := zerolog.ParseLevel(config.LogLevel)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse log level")
	}
	log.Logger = log.Logger.Level(logLvl).With().Timestamp().Logger()
}

func main() {
	log.Info().Msg("Sup")
	mainClient := resty.New()

	weather, err := GetWeather(mainClient, config.Latitude, config.Longitude, time.Now().AddDate(0, 0, -3), time.Now().AddDate(0, 0, -3))
	if err != nil {
		log.Error().Err(err).Msg("Failed to get weather data")
		return
	}
	log.Info().Interface("Weather", weather).Msg("Got weather data")
}
