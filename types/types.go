package types

import "net/url"

var WeatherApiDateFormat = "2006-01-02"

type GeolocationData struct {
	Latitude	float64 `url:"latitude" validate:"required,latitude"`
	Longitude	float64 `url:"longitude" validate:"required,longitude"`
}

type SetParamsFunc func(geoData GeolocationData) url.Values

// source: https://www.nodc.noaa.gov/archive/arc0021/0002199/1.1/data/0-data/HTML/WMO-CODE/WMO4677.HTM
var RainCodes = map[int]DominantWeather{
	21: WithRainfall,
	23: WithRainfall,
	24: WithRainfall, // can be either snow or rain
	25: WithRainfall,
	26: WithRainfall, // can be either snow or rain
	27: WithRainfall, // can be either snow or rain
	58: WithRainfall,
	59: WithRainfall,
	60: WithRainfall,
	61: WithRainfall,
	63: WithRainfall,
	65: WithRainfall,
	66: WithRainfall,
	67: WithRainfall,
	68: WithRainfall,
	69: WithRainfall,
	80: WithRainfall,
	81: WithRainfall,
	82: WithRainfall,
	83: WithRainfall,
	84: WithRainfall,
	87: WithRainfall,
	88: WithRainfall,
	89: WithRainfall,
	90: WithRainfall,
	91: WithRainfall,
	92: WithRainfall,
	93: WithRainfall,
	94: WithRainfall,
	95: WithRainfall,
	97: WithRainfall,
}
