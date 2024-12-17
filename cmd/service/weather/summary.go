package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"time"

	"github.com/DroppedHard/weather-forecast-api/types"
	"gonum.org/v1/gonum/floats"
)

var summaryParams = map[string]string{
	"daily":      	"temperature_2m_max,temperature_2m_min,apparent_temperature_max,apparent_temperature_min,daylight_duration,sunshine_duration,weather_code",
	"hourly": 		"surface_pressure",
	"timezone":   	"GMT", // if time allows - get this from frontend to get user's timezone data
}


func setSummaryRequestParams(geoData types.GeolocationData) (params url.Values) {
	currentDate := time.Now()
	futureDate := currentDate.Add(7 * 24 * time.Hour)
	params = url.Values{}
	params.Add("latitude", strconv.FormatFloat(geoData.Latitude, 'f', -1, 64))
	params.Add("longitude", strconv.FormatFloat(geoData.Longitude, 'f', -1, 64))
	params.Add("startDate", currentDate.Format(types.WeatherApiDateFormat))
	params.Add("endDate", futureDate.Format(types.WeatherApiDateFormat))
	for key, value := range summaryParams {
		params.Add(key, value)
	}
	return
}

func processSummaryData(body io.ReadCloser) (types.SummaryResponseData, error){
	var apiResponse types.WeatherApiResponse

	if err := json.NewDecoder(body).Decode(&apiResponse); err != nil {
		return types.SummaryResponseData{}, fmt.Errorf("failed to decode response: %v", err)
	}
	defer body.Close()
	
	forecastResponse := types.SummaryResponseData{}
	data, err := json.Marshal(apiResponse)
	if err != nil {
		return types.SummaryResponseData{}, fmt.Errorf("failed to marshal weather API response: %v", err)
	}
	if err := json.Unmarshal(data, &forecastResponse); err != nil {
		return types.SummaryResponseData{}, fmt.Errorf("failed to unmarshal into forecast response: %v", err)
	}
	forecastResponse.AvgDaylightDuration = prepareSummaryData(average(apiResponse.Daily.DaylightDuration), apiResponse.DailyUnits.DaylightDuration)
	forecastResponse.AvgSurfacePressure = prepareSummaryData(average(apiResponse.Hourly.SurfacePressure), apiResponse.HourlyUnits.SurfacePressure)
	forecastResponse.MinTemperature = prepareSummaryData(floats.Min(apiResponse.Daily.TemperatureMin), apiResponse.HourlyUnits.SurfacePressure)
	forecastResponse.MinApparentTemperature = prepareSummaryData(floats.Min(apiResponse.Daily.ApparentTemperatureMax), apiResponse.HourlyUnits.SurfacePressure)
	forecastResponse.MaxTemperature = prepareSummaryData(floats.Max(apiResponse.Daily.TemperatureMax), apiResponse.HourlyUnits.SurfacePressure)
	forecastResponse.MaxApparentTemperature = prepareSummaryData(floats.Max(apiResponse.Daily.ApparentTemperatureMax), apiResponse.HourlyUnits.SurfacePressure)
	forecastResponse.DominantWeather = calculateDominantWeather(apiResponse.Daily.WeatherCode)
	return forecastResponse, nil
}

func average(xs []float64) float64 {
	return floats.Sum(xs)/float64(len(xs))
}

func prepareSummaryData(value float64, unit string) string {
	return strconv.FormatFloat(value, 'f', -1, 64) + " "+ unit
}



func calculateDominantWeather(weatherCodes []int) types.DominantWeather {
	counter := 0
	for _, code := range weatherCodes {
		if types.RainCodes[code] {
			counter++
		}
	}
	if counter > 3 {
		return types.WithRainfall
	}
	return types.WithoutRainfall
}