package types

type GeolocationData struct {
	Latitude	float64 `url:"latitude" validate:"required,latitude"`
	Longitude	float64 `url:"longitude" validate:"required,longitude"`
}

type SharedWeatherData struct {
	Latitude            float64 `json:"latitude"`
	Longitude           float64 `json:"longitude"`
	GenerationTime    float64 `json:"generationtime_ms"`
	Timezone            string  `json:"timezone"`
	TimezoneAbbr        string  `json:"timezone_abbreviation"`
	Elevation           float64 `json:"elevation"`
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

type WeatherApiResponse struct {
	SharedWeatherData
	DailyUnits          DailyUnits `json:"daily_units"`
	Daily               DailyData  `json:"daily"`
}

type DailyUnits struct {
	SharedDailyUnits
	DaylightDuration string `json:"daylight_duration"`
}

type DailyData struct {
	SharedDailyData
	DaylightDuration   []float64 `json:"daylight_duration"`
}

type ForecastResponseData struct {
	SharedWeatherData
	ParsedDailyUnits   	ParsedDailyUnits `json:"daily_units"`
	ParsedDaily         ParsedDailyData  `json:"daily"`
}

type ParsedDailyUnits struct {
	SharedDailyUnits
	EstimatedEnergyGenerated	string `json:"estimated_energy_generated"`
}

type ParsedDailyData struct {
	SharedDailyData
	EstimatedEnergyGenerated   	[]float64 `json:"estimated_energy_generated"`
}