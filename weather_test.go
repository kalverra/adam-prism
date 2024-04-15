package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"github.com/test-go/testify/assert"
	"github.com/test-go/testify/require"
)

func TestGetWeather(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		mock := `{"daily":{"apparent_temperature_max":[27.3],"apparent_temperature_mean":[18.9],"apparent_temperature_min":[12.6],"daylight_duration":[53343.98],"precipitation_sum":[0.094],"rain_sum":[0],"snowfall_sum":[0.661],"sunrise":["2024-04-11T15:10"],"sunset":["2024-04-12T05:55"],"sunshine_duration":[6931.08],"temperature_2m_max":[35.7],"temperature_2m_mean":[26.5],"temperature_2m_min":[20.1],"time":["2024-04-12"],"weather_code":[73],"wind_speed_10m_max":[9.6]},"daily_units":{"apparent_temperature_max":"°F","apparent_temperature_mean":"°F","apparent_temperature_min":"°F","daylight_duration":"s","precipitation_sum":"inch","rain_sum":"inch","snowfall_sum":"inch","sunrise":"iso8601","sunset":"iso8601","sunshine_duration":"s","temperature_2m_max":"°F","temperature_2m_mean":"°F","temperature_2m_min":"°F","time":"iso8601","weather_code":"wmo code","wind_speed_10m_max":"mp/h"},"elevation":585,"generationtime_ms":0.09500980377197266,"latitude":64.077324,"longitude":141.87668,"timezone":"America/New_York","timezone_abbreviation":"EDT","utc_offset_seconds":-14400}`
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		_, err := w.Write([]byte(mock))
		assert.NoError(t, err)
		log.Trace().Str("Data", mock).Msg("Mock response sent")
	}))
	t.Cleanup(server.Close)

	client := resty.New()
	client.SetBaseURL(server.URL)
	data, err := GetWeather(client, 0, 0, time.Now(), time.Now())
	require.NoError(t, err)
	require.NotNil(t, data)
	require.NotEmpty(t, data.Daily.ApparentTemperatureMax, "ApparentTemperatureMax should not be empty %+v", data)
	require.Equal(t, 27.3, data.Daily.ApparentTemperatureMax[0])
}

func TestGetWeatherError(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(500)
	}))
	t.Cleanup(server.Close)

	client := resty.New()
	client.SetBaseURL(server.URL)
	data, err := GetWeather(client, 0, 0, time.Now(), time.Now())
	require.Error(t, err)
	require.Nil(t, data)
}
