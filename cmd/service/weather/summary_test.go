package weather

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	"github.com/DroppedHard/weather-forecast-api/types"
	"github.com/stretchr/testify/assert"
)

func TestSetSummaryRequestParams(t *testing.T) {
	geoData := types.GeolocationData{
		Latitude:  52.52,
		Longitude: 13.4050,
	}

	params := setSummaryRequestParams(geoData)

	assert.Equal(t, params.Get("latitude"), "52.52")
	assert.Equal(t, params.Get("longitude"), "13.405")
	assert.NotEmpty(t, params.Get("startDate"))
	assert.NotEmpty(t, params.Get("endDate"))
	assert.Contains(t, params, "daily")
	assert.Contains(t, params, "timezone")
}

func TestProcessSummaryData(t *testing.T) {
	mockApiResponse := types.WeatherApiResponse{
		Daily: types.DailyData{
			SharedDailyData: types.SharedDailyData{
				WeatherCode:      []int{21, 21, 21, 21}, // it's all with rainfall
				TemperatureMax:   []float64{25.0, 28.0, 30.0, 27.5},
				TemperatureMin:   []float64{15.0, 18.0, 16.5, 17.0},
				Time: []string{"2024-12-18", "2024-12-19", "2024-12-20", "2024-12-21"},
			},
			ApparentTemperatureMax: []float64{24.5, 27.5, 29.0, 26.0},
			ApparentTemperatureMin: []float64{14.5, 17.0, 15.5, 16.0},
			DaylightDuration: []float64{36000, 35000, 37000, 35500},
		},
		Hourly: types.HourlyData{
			SurfacePressure: []float64{1015.5, 1014.0, 1016.2, 1015.7},
		},
		DailyUnits: types.DailyUnits{
			SharedDailyUnits: types.SharedDailyUnits{
				TemperatureMin: "°C",
				TemperatureMax: "°C",
			},
			DaylightDuration: "s",
			ApparentTemperatureMax: "°C",
			ApparentTemperatureMin: "°C",
		},
		HourlyUnits: types.HourlyUnits{
			SurfacePressure: "hPa",
		},
	}

	mockData, err := json.Marshal(mockApiResponse)
	if err != nil {
		t.Fatalf("Failed to marshal mock API response: %v", err)
	}

	body := io.NopCloser(bytes.NewReader(mockData))

	forecastResponse, err := processSummaryData(body)

	assert.NoError(t, err)
	assert.NotNil(t, forecastResponse)
	assert.Equal(t, forecastResponse.AvgDaylightDuration, "35875 s")
	assert.Equal(t, forecastResponse.AvgSurfacePressure, "1015.35 hPa")
	assert.Equal(t, forecastResponse.MinTemperature, "15 °C")
	assert.Equal(t, forecastResponse.MinApparentTemperature, "14.5 °C")
	assert.Equal(t, forecastResponse.MaxTemperature, "30 °C")
	assert.Equal(t, forecastResponse.MaxApparentTemperature, "29 °C")
	assert.Equal(t, forecastResponse.DominantWeather, types.WithRainfall)
}

func TestAverageFunction(t *testing.T) {
	data := []float64{36000, 35000, 37000, 35500}

	avg := average(data)

	assert.Equal(t, avg, 35875.0)
}

func TestPrepareSummaryData(t *testing.T) {
	value := 25.5
	unit := "°C"
	expected := "25.5 °C"

	result := prepareSummaryData(value, unit)

	assert.Equal(t, result, expected)
}

func TestCalculateDominantWeather(t *testing.T) {
	weatherCodes := []int{1, 21, 21, 21, 1, 21, 1}
	dominantWeather := calculateDominantWeather(weatherCodes)

	// Assert that the dominant weather is "With Rainfall"
	assert.Equal(t, dominantWeather, types.WithRainfall)

	// Test case without enough rain codes
	weatherCodes2 := []int{1, 1, 1, 1, 21, 21, 1}
	dominantWeather2 := calculateDominantWeather(weatherCodes2)

	// Assert that the dominant weather is "Without Rainfall"
	assert.Equal(t, dominantWeather2, types.WithoutRainfall)
}
