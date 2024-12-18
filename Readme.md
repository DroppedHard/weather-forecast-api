# Weather-Forecast-API

This is a GoLang service that is a somehow fascade to [Open Meteo API](https://open-meteo.com) that has 2 simple endpoints to get weekly forecast and weekly summary. 

# How to run
You can use makefile to simply run the app. 
1. [install GoLang](https://go.dev/doc/install) 
2. Clone the repository
3. Build it
```
make build # or go build -o bin/weather-forecast-api cmd/main.go
```
3. Go for it
```
make run # or ./bin/weather-forecast-api - if it was executed
```


## Windows users

To use MakeFile as a Windows user - [install Chocolatey](https://chocolatey.org) and install chocolatey:
```cmd
choco install make
```

# Endpoints
1. `/api/v1/forecast`
- Receives `Longitude` and `Latitude` as parameters (case sensitive!)
- Returns weekly forecast from current day. Example return:

`http://localhost:8081/api/v1/forecast?Latitude=52.52&Longitude=13.41`
```json
{
    "latitude": 52.52,
    "longitude": 13.419998,
    "generationtime_ms": 0.11599063873291016,
    "elevation": 38,
    "timezone": "GMT",
    "timezone_abbreviation": "GMT",
    "daily_units": {
        "time": "iso8601",
        "weather_code": "wmo code",
        "temperature_2m_max": "°C",
        "temperature_2m_min": "°C",
        "estimated_energy_generated": "kWh"
    },
    "daily": {
        "time": [
            "2024-12-18",
            "2024-12-19",
            "2024-12-20",
            "2024-12-21",
            "2024-12-22",
            "2024-12-23",
            "2024-12-24"
        ],
        "weather_code": [
            61,
            61,
            80,
            61,
            85,
            85,
            61
        ],
        "temperature_2m_max": [
            11.6,
            12.2,
            5.9,
            6,
            6.3,
            6.3,
            6.2
        ],
        "temperature_2m_min": [
            5.9,
            4.8,
            3.1,
            1.6,
            2.6,
            2.5,
            2.8
        ],
        "estimated_energy_generated": [
            3.83,
            3.83,
            3.82,
            3.82,
            3.83,
            3.83,
            3.83
        ]
    }
}
```

2. `/api/v1/summary`
- Receives `Longitude` and `Latitude` as parameters (case sensitive!)
- Returns weekly summary from current day. Example return:

`http://localhost:8081/api/v1/summary?Latitude=52.52&Longitude=13.41`
```json
{
    "latitude": 52.52,
    "longitude": 13.419998,
    "generationtime_ms": 0.26297569274902344,
    "elevation": 38,
    "timezone": "GMT",
    "timezone_abbreviation": "GMT",
    "avg_surface_pressure": "1003.14 hPa",
    "avg_daylight_duration": "27551.71 s",
    "min_temperature": "1.6 °C",
    "max_temperature": "12.2 °C",
    "min_apparent_temperature": "-2.3 °C",
    "max_apparent_temperature": "8.7 °C",
    "dominant_weather": true
}
```