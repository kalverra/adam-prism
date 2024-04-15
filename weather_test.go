package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/test-go/testify/assert"
	"github.com/test-go/testify/require"
)

func TestGetWeather(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		mock := `{"latitude":64.077324,"longitude":141.87668,"generationtime_ms":0.42998790740966797,"utc_offset_seconds":-14400,"timezone":"America/New_York","timezone_abbreviation":"EDT","elevation":585.0,"daily_units":{"time":"iso8601","weather_code":"wmo code","temperature_2m_max":"°F","temperature_2m_min":"°F","temperature_2m_mean":"°F","apparent_temperature_max":"°F","apparent_temperature_min":"°F","apparent_temperature_mean":"°F","sunrise":"iso8601","sunset":"iso8601","daylight_duration":"s","sunshine_duration":"s","precipitation_sum":"inch","rain_sum":"inch","snowfall_sum":"inch","precipitation_hours":"h","wind_speed_10m_max":"mp/h"},"daily":{"time":["2024-04-12"],"weather_code":[73],"temperature_2m_max":[35.7],"temperature_2m_min":[20.1],"temperature_2m_mean":[26.5],"apparent_temperature_max":[27.3],"apparent_temperature_min":[12.6],"apparent_temperature_mean":[18.9],"sunrise":["2024-04-11T15:10"],"sunset":["2024-04-12T05:55"],"daylight_duration":[53343.98],"sunshine_duration":[6931.08],"precipitation_sum":[0.094],"rain_sum":[0.000],"snowfall_sum":[0.661],"precipitation_hours":[11.0],"wind_speed_10m_max":[9.6]}}`
		w.WriteHeader(200)
		_, err := w.Write([]byte(mock))
		assert.NoError(t, err)
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
