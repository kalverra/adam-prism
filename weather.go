package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
)

// Gathering historical weather data from Open-Meteo API: https://open-meteo.com/

// GetWeather retrieves historical weather data from the Open-Meteo API for a given location and date range
func GetWeather(client *resty.Client, latitude float64, longitude float64, startDate time.Time, endDate time.Time) (*WeatherData, error) {
	log.Debug().Msg("Getting Weather Data")
	daily := []string{"weather_code", "temperature_2m_max", "temperature_2m_min", "temperature_2m_mean", "apparent_temperature_max", "apparent_temperature_min", "apparent_temperature_mean", "sunrise", "sunset", "daylight_duration", "sunshine_duration", "precipitation_sum", "rain_sum", "snowfall_sum", "precipitation_hours", "wind_speed_10m_max"}
	queryParams := map[string]string{
		"latitude":           fmt.Sprintf("%f", latitude),
		"longitude":          fmt.Sprintf("%f", longitude),
		"daily":              strings.Join(daily, ","),
		"temperature_unit":   "fahrenheit",
		"wind_speed_unit":    "mph",
		"precipitation_unit": "inch",
		"timezone":           "America/New_York",
		"start_date":         startDate.Format("2006-01-02"),
		"end_date":           endDate.Format("2006-01-02"),
	}

	result := &WeatherData{}
	if !strings.Contains(client.BaseURL, "127.0.0.1") { // In case of testing
		client.SetBaseURL("https://archive-api.open-meteo.com")
	}
	resp, err := client.R().
		SetQueryParams(queryParams).
		SetHeader("Accept", "application/json").
		SetResult(result).
		Get("/v1/archive")
	if err != nil || resp.IsError() {
		return nil, fmt.Errorf("error getting weather data with status code %d: %w", resp.StatusCode(), err)
	}

	log.Debug().Str("Data", resp.String()).Interface("Marshalled", result).Msg("Got Weather Data")
	return result, nil
}

// WeatherData holds response data from the Open-Meteo API
type WeatherData struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	GenerationtimeMs     float64 `json:"generationtime_ms"`
	UtcOffsetSeconds     int     `json:"utc_offset_seconds"`
	Timezone             string  `json:"timezone"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation"`
	Elevation            float64 `json:"elevation"`
	DailyUnits           struct {
		Time                    string `json:"time"`
		WeatherCode             string `json:"weather_code"`
		Temperature2MMax        string `json:"temperature_2m_max"`
		Temperature2MMin        string `json:"temperature_2m_min"`
		Temperature2MMean       string `json:"temperature_2m_mean"`
		ApparentTemperatureMax  string `json:"apparent_temperature_max"`
		ApparentTemperatureMin  string `json:"apparent_temperature_min"`
		ApparentTemperatureMean string `json:"apparent_temperature_mean"`
		Sunrise                 string `json:"sunrise"`
		Sunset                  string `json:"sunset"`
		DaylightDuration        string `json:"daylight_duration"`
		SunshineDuration        string `json:"sunshine_duration"`
		PrecipitationSum        string `json:"precipitation_sum"`
		RainSum                 string `json:"rain_sum"`
		SnowfallSum             string `json:"snowfall_sum"`
		WindSpeed10MMax         string `json:"wind_speed_10m_max"`
	} `json:"daily_units"`
	Daily struct {
		Time                    []string  `json:"time"`
		WeatherCode             []int     `json:"weather_code"`
		Temperature2MMax        []float64 `json:"temperature_2m_max"`
		Temperature2MMin        []float64 `json:"temperature_2m_min"`
		Temperature2MMean       []any     `json:"temperature_2m_mean"`
		ApparentTemperatureMax  []float64 `json:"apparent_temperature_max"`
		ApparentTemperatureMin  []float64 `json:"apparent_temperature_min"`
		ApparentTemperatureMean []any     `json:"apparent_temperature_mean"`
		Sunrise                 []string  `json:"sunrise"`
		Sunset                  []string  `json:"sunset"`
		DaylightDuration        []float64 `json:"daylight_duration"`
		SunshineDuration        []any     `json:"sunshine_duration"`
		PrecipitationSum        []any     `json:"precipitation_sum"`
		RainSum                 []any     `json:"rain_sum"`
		SnowfallSum             []any     `json:"snowfall_sum"`
		WindSpeed10MMax         []float64 `json:"wind_speed_10m_max"`
	} `json:"daily"`
}
