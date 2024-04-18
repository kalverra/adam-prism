package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/test-go/testify/require"
)

var mockWeatherData = &WeatherData{
	Latitude:             64.077324,
	Longitude:            141.87668,
	Timezone:             "America/New_York",
	TimezoneAbbreviation: "EDT",
	UtcOffsetSeconds:     -14400,
	Elevation:            585,
	GenerationtimeMs:     0.09500980377197266,
	Daily: WeatherDaily{
		Time:                    []string{"2024-04-12"},
		ApparentTemperatureMax:  []float64{27.3},
		ApparentTemperatureMean: []float64{18.9},
		ApparentTemperatureMin:  []float64{12.6},
		DaylightDuration:        []float64{53343.98},
		PrecipitationSum:        []float64{0.094},
		RainSum:                 []float64{0},
		SnowfallSum:             []float64{0.661},
		Sunrise:                 []string{"2024-04-11T15:10"},
		Sunset:                  []string{"2024-04-12T05:55"},
		SunshineDuration:        []float64{6931.08},
		WindSpeed10MMax:         []float64{9.6},
	},
}

func TestAddWeatherData(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(200)
	}))
	client := resty.New()
	client.SetBaseURL(server.URL)

	err := AddWeatherData(client, mockWeatherData)
	require.NoError(t, err)
}

func TestAddWeatherDataError(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(500)
	}))
	client := resty.New()
	client.SetBaseURL(server.URL)

	err := AddWeatherData(client, mockWeatherData)
	require.Error(t, err, "Expecting error when status code is not 200")
}
