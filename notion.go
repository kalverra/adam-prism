package main

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
)

type NotionPageData struct {
	Parent     map[string]string      `json:"parent"`
	Properties map[string]interface{} `json:"properties"`
}

// AddWeatherData adds 1 to many weather data records to the WeatherDB database
func AddWeatherData(client *resty.Client, weatherData *WeatherData) error {
	if !strings.Contains(client.BaseURL, "127.0.0.1") { // For testing
		client.SetBaseURL("https://api.notion.com/v1/")
	}

	for _, data := range getNotionWeatherData(weatherData) {
		log.Debug().Interface("Data", data).Msg("Adding Weather Data to Notion")
		// https://developers.notion.com/reference/post-page
		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetHeader("Authorization", fmt.Sprintf("Bearer %s", config.NotionKey)).
			SetHeader("Notion-Version", "2022-06-28"). // Ensure this version is up to date
			SetBody(data).
			Post("pages")
		if err != nil {
			return fmt.Errorf("error adding weather data to notion with status code %d: %w", resp.StatusCode(), err)
		}
		if resp.IsError() {
			return fmt.Errorf("error adding weather data to notion with status code %d: %s", resp.StatusCode(), resp.String())
		}
	}

	return nil
}

// getNotionWeatherData converts WeatherData to NotionPageData for the Notion API
func getNotionWeatherData(weatherData *WeatherData) []*NotionPageData {
	if weatherData == nil {
		return []*NotionPageData{}
	}
	notionData := []*NotionPageData{}
	for index, weatherDataDate := range weatherData.Daily.Time {
		nd := &NotionPageData{
			Parent: map[string]string{
				"database_id": config.WeatherDB,
				"type":        "database_id",
			},
			Properties: map[string]any{
				"Date": map[string]any{
					"type": "date",
					"date": map[string]string{"start": weatherDataDate},
				},
				"Latitude": map[string]any{
					"type":   "number",
					"number": weatherData.Latitude,
				},
				"Longitude": map[string]any{
					"type":   "number",
					"number": weatherData.Longitude,
				},
				"Min Temp": map[string]any{
					"type":   "number",
					"number": weatherData.Daily.ApparentTemperatureMin[index],
				},
				"Max Temp": map[string]any{
					"type":   "number",
					"number": weatherData.Daily.ApparentTemperatureMax[index],
				},
				"Mean Temp": map[string]any{
					"type":   "number",
					"number": weatherData.Daily.ApparentTemperatureMean[index],
				},
				"Sunrise": map[string]any{
					"type": "date",
					"date": map[string]any{"start": weatherData.Daily.Sunrise[index]},
				},
				"Sunset": map[string]any{
					"type": "date",
					"date": map[string]any{"start": weatherData.Daily.Sunset[index]},
				},
				"Daylight Duration": map[string]any{
					"type":   "number",
					"number": weatherData.Daily.DaylightDuration[index],
				},
				"Sunshine Duration": map[string]any{
					"type":   "number",
					"number": weatherData.Daily.SunshineDuration[index],
				},
				"Precipitation Sum": map[string]any{
					"type":   "number",
					"number": weatherData.Daily.PrecipitationSum[index],
				},
				"Rain Sum": map[string]any{
					"type":   "number",
					"number": weatherData.Daily.RainSum[index],
				},
				"Snowfall Sum": map[string]any{
					"type":   "number",
					"number": weatherData.Daily.SnowfallSum[index],
				},
				"Wind Speed Max": map[string]any{
					"type":   "number",
					"number": weatherData.Daily.WindSpeed10MMax[index],
				},
			},
		}
		notionData = append(notionData, nd)
	}
	return notionData
}
