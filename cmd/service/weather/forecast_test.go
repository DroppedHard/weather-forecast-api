package weather

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"
	"time"

	"github.com/DroppedHard/weather-forecast-api/types"
	"github.com/stretchr/testify/assert"
)

func TestSetForecastRequestParams(t *testing.T) {
	geoData := types.GeolocationData{
		Latitude:  52.52,
		Longitude: 13.4050,
	}

	params := setForecastRequestParams(geoData)

	assert.Equal(t, "52.52", params.Get("latitude"), "Expected latitude parameter to be 52.52")
	assert.Equal(t, "13.405", params.Get("longitude"), "Expected longitude parameter to be 13.405")
	assert.Equal(t, time.Now().Format(types.WeatherApiDateFormat), params.Get("startDate"), "Expected correct startDate")
	assert.Equal(t, time.Now().Add(7*24*time.Hour).Format(types.WeatherApiDateFormat), params.Get("endDate"), "Expected correct endDate")
	assert.Equal(t, "weather_code,temperature_2m_max,temperature_2m_min,daylight_duration", params.Get("daily"), "Expected correct daily parameter")
	assert.Equal(t, "GMT", params.Get("timezone"), "Expected timezone to be GMT")
}

func TestProcessWeatherData_Success(t *testing.T) {
	mockApiResponse := types.WeatherApiResponse{
        Daily: types.DailyData{
            SharedDailyData: types.SharedDailyData{
                WeatherCode:      []int{1, 2},
                TemperatureMax:   []float64{20.5, 22.3},
                TemperatureMin:   []float64{10.0, 11.5},
                Time:             []string{"2024-12-18", "2024-12-19"}, 
            },
            DaylightDuration: []float64{36000, 35000},
            ApparentTemperatureMax: []float64{25.0, 26.5},
            ApparentTemperatureMin: []float64{15.0, 16.0},
        },
    }

	body, err := json.Marshal(mockApiResponse)
	assert.NoError(t, err)

	bodyReader := io.NopCloser(bytes.NewReader(body))

	result, err := processWeatherData(bodyReader)
	assert.NoError(t, err)

	assert.Equal(t, "kWh", result.ParsedDailyUnits.EstimatedEnergyGenerated, "Expected energy unit to be 'kWh'")

	expectedEnergy := []float64{2.5 * 0.2 * 10, 2.5 * 0.2 * 9.72} // 10 hours, 9.72 hours
	assert.Equal(t, expectedEnergy[0], result.ParsedDaily.EstimatedEnergyGenerated[0], "Expected correct energy estimate for first day")
	assert.Equal(t, expectedEnergy[1], result.ParsedDaily.EstimatedEnergyGenerated[1], "Expected correct energy estimate for second day")
}

func TestProcessWeatherData_Error(t *testing.T) {
	malformedBody := io.NopCloser(bytes.NewReader([]byte("malformed json")))

	result, err := processWeatherData(malformedBody)

	assert.Error(t, err)
	assert.Equal(t, types.ForecastResponseData{}, result, "Expected empty result on error")
}

func TestCalculateEnergyEstimate(t *testing.T) {
	daylightDuration := []float64{36000, 35000}
	expectedEnergy := []float64{2.5 * 0.2 * 10, 2.5 * 0.2 * 9.72}

	result := calculateEnergyEstimate(daylightDuration)

	assert.Equal(t, expectedEnergy[0], result[0], "Expected correct energy estimate for first day")
	assert.Equal(t, expectedEnergy[1], result[1], "Expected correct energy estimate for second day")
}

func TestCalculateEnergyEstimate_EmptyInput(t *testing.T) {
	daylightDuration := []float64{}

	result := calculateEnergyEstimate(daylightDuration)

	assert.Empty(t, result, "Expected empty result for empty input")
}
