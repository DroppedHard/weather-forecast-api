package types

type WeatherApiResponse struct {
	SharedWeatherData
	HourlyUnits         HourlyUnits `json:"hourly_units"`
	Hourly              HourlyData  `json:"hourly"`
	DailyUnits          DailyUnits 	`json:"daily_units"`
	Daily               DailyData  	`json:"daily"`
}

type DailyUnits struct {
	SharedDailyUnits
	DaylightDuration 			string 	`json:"daylight_duration"`
	ApparentTemperatureMax     	string  `json:"apparent_temperature_max"`
	ApparentTemperatureMin     	string  `json:"apparent_temperature_min"`
}

type DailyData struct {
	SharedDailyData
	DaylightDuration   			[]float64 `json:"daylight_duration"`
	ApparentTemperatureMax      []float64 `json:"apparent_temperature_max"`
	ApparentTemperatureMin      []float64 `json:"apparent_temperature_min"`
}

type HourlyUnits struct {
	Time            string `json:"time"`
	SurfacePressure string `json:"surface_pressure"`
}

type HourlyData struct {
	Time            []string  `json:"time"`
	SurfacePressure []float64 `json:"surface_pressure"`
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

type SummaryResponseData struct {
	SharedWeatherData
	AvgSurfacePressure 			string 			`json:"avg_surface_pressure"`
	AvgDaylightDuration 		string 			`json:"avg_daylight_duration"`
	MinTemperature				string 			`json:"min_temperature"`
	MaxTemperature				string 			`json:"max_temperature"`
	MinApparentTemperature		string 			`json:"min_apparent_temperature"`
	MaxApparentTemperature		string 			`json:"max_apparent_temperature"`
	DominantWeather				DominantWeather `json:"dominant_weather"`
}

type DominantWeather bool

const (
	WithRainfall  	DominantWeather	= true
	WithoutRainfall DominantWeather = false
)