package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/DroppedHard/weather-forecast-api/config"
	"github.com/DroppedHard/weather-forecast-api/types"
	"github.com/DroppedHard/weather-forecast-api/utils"
)

func GetWeatherForecast(geoData types.GeolocationData) (*http.Response, error) {
	weatherUrl, err := url.Parse(config.Envs.WeatherApi + "forecast")
	if err != nil {
		return nil, fmt.Errorf("failed to parse weather API URL: %w", err)
	}
	params := utils.SetRequestParams(geoData)
	weatherUrl.RawQuery = params.Encode()

	weatherReq, err := http.NewRequest("GET", weatherUrl.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create weather API request: %w", err)
	}
	weatherReq.Header.Add("Accept", "application/json")

	weatherResp, err := http.DefaultClient.Do(weatherReq)
	
	if err != nil {
		return nil, fmt.Errorf("failed to execute weather API request: %w", err)
	}
	if weatherResp.StatusCode != http.StatusOK {
		defer weatherResp.Body.Close()
		body, _ := io.ReadAll(weatherResp.Body)
		return nil, fmt.Errorf("weather API error: (%v): %s", weatherResp.StatusCode, body)
	}
	return weatherResp, nil
}

func ProcessWeatherData(body io.ReadCloser) (types.ForecastResponseData, error){
	var apiResponse types.WeatherApiResponse

	if err := json.NewDecoder(body).Decode(&apiResponse); err != nil {
		return types.ForecastResponseData{}, fmt.Errorf("failed to decode response: %v", err)
	}
	defer body.Close()
	estimatedEnergy := CalculateEnergyEstimate(apiResponse.Daily.DaylightDuration)
	
	forecastResponse := types.ForecastResponseData{}
	data, err := json.Marshal(apiResponse)
	if err != nil {
		return types.ForecastResponseData{}, fmt.Errorf("failed to marshal weather API response: %v", err)
	}
	if err := json.Unmarshal(data, &forecastResponse); err != nil {
		return types.ForecastResponseData{}, fmt.Errorf("failed to unmarshal into forecast response: %v", err)
	}
	forecastResponse.ParsedDailyUnits.EstimatedEnergyGenerated = "kWh"
	forecastResponse.ParsedDaily.EstimatedEnergyGenerated = estimatedEnergy
	return forecastResponse, nil
}

func CalculateEnergyEstimate(daylightDuration []float64) (estimatedEnergy []float64) {
	estimatedEnergy = make([]float64, len(daylightDuration))
	for i, daylightSeconds := range daylightDuration {
		// Calculate energy: 2.5kW * 0.2 efficiency * daylight duration in hours
		daylightHours := daylightSeconds / 3600.0
		estimatedEnergy[i] = 2.5 * 0.2 * daylightHours
	}
	return
}