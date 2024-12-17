package types

type SharedWeatherData struct {
	Latitude            float64 `json:"latitude"`
	Longitude           float64 `json:"longitude"`
	GenerationTime    	float64 `json:"generationtime_ms"`
	Elevation           float64 `json:"elevation"`
	Timezone            string  `json:"timezone"`
	TimezoneAbbr        string  `json:"timezone_abbreviation"`
}

type SharedDailyUnits struct {
	Time             string `json:"time"`
	WeatherCode      string `json:"weather_code"`
	TemperatureMax   string `json:"temperature_2m_max"`
	TemperatureMin   string `json:"temperature_2m_min"`
}

type SharedDailyData struct {
	Time               []string  `json:"time"`
	WeatherCode        []int     `json:"weather_code"`
	TemperatureMax     []float64 `json:"temperature_2m_max"`
	TemperatureMin     []float64 `json:"temperature_2m_min"`
}