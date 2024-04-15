package main

import (
	"os"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var config *Config

// read config and setup logger
func init() {
	var err error
	zerolog.TimeFieldFormat = time.RFC3339Nano
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "15:04:05.00", // hh:mm:ss.ss format
	})

	config, err = ReadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read config")
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

	_, err := GetWeather(mainClient, config.Latitude, config.Longitude, time.Now().AddDate(0, 0, -3), time.Now().AddDate(0, 0, -3))
	if err != nil {
		log.Error().Err(err).Msg("Failed to get weather data")
		return
	}
}
