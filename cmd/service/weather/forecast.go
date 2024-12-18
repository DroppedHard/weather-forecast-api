package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"time"

	"github.com/DroppedHard/weather-forecast-api/types"
	"github.com/DroppedHard/weather-forecast-api/utils"
)

var forecastParams = map[string]string{
	"daily":      "weather_code,temperature_2m_max,temperature_2m_min,daylight_duration",
	"timezone":   "GMT",// if time allows - get this from frontend to get user's timezone data
}


func setForecastRequestParams(geoData types.GeolocationData) (params url.Values) {
	currentDate := time.Now()
	futureDate := currentDate.Add(7 * 24 * time.Hour)
	params = url.Values{}
	params.Add("latitude", strconv.FormatFloat(geoData.Latitude, 'f', -1, 64))
	params.Add("longitude", strconv.FormatFloat(geoData.Longitude, 'f', -1, 64))
	params.Add("startDate", currentDate.Format(types.WeatherApiDateFormat))
	params.Add("endDate", futureDate.Format(types.WeatherApiDateFormat))
	for key, value := range forecastParams {
		params.Add(key, value)
	}
	return
}

func processWeatherData(body io.ReadCloser) (types.ForecastResponseData, error){
	var apiResponse types.WeatherApiResponse

	if err := json.NewDecoder(body).Decode(&apiResponse); err != nil {
		return types.ForecastResponseData{}, fmt.Errorf("failed to decode response: %v", err)
	}
	defer body.Close()
	
	forecastResponse := types.ForecastResponseData{}
	data, err := json.Marshal(apiResponse)
	if err != nil {
		return types.ForecastResponseData{}, fmt.Errorf("failed to marshal weather API response: %v", err)
	}
	if err := json.Unmarshal(data, &forecastResponse); err != nil {
		return types.ForecastResponseData{}, fmt.Errorf("failed to unmarshal into forecast response: %v", err)
	}
	forecastResponse.ParsedDailyUnits.EstimatedEnergyGenerated = "kWh"
	forecastResponse.ParsedDaily.EstimatedEnergyGenerated = calculateEnergyEstimate(apiResponse.Daily.DaylightDuration)
	return forecastResponse, nil
}

func calculateEnergyEstimate(daylightDuration []float64) (estimatedEnergy []float64) {
	estimatedEnergy = make([]float64, len(daylightDuration))
	for i, daylightSeconds := range daylightDuration {
		// Calculate energy: 2.5kW * 0.2 efficiency * daylight duration in hours
		daylightHours := daylightSeconds / 3600.0
		estimatedEnergy[i] = utils.RoundFloat(2.5 * 0.2 * daylightHours, 2)
	}
	return 
}