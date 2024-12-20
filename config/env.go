package config

import (
	"log"
	"os"

	"github.com/lpernett/godotenv"
)

type Config struct {
	PORT			string
	WEATHER_API		string
	ORIGINS_ALLOWED	string
}

var Envs = initConfig()

func initConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Missing .env file - working with included env variables...")
	}
	
	return Config{
		PORT: getEnv("PORT", "8080"),
		WEATHER_API: getEnv("WEATHER_API", "https://api.open-meteo.com/v1/"),
		ORIGINS_ALLOWED: getEnv("ORIGINS_ALLOWED", `https://pogodynka-app.netlify.app,https://weather-forecast-app-j7gf.onrender.com,http://localhost:4173,http://localhost:5173`),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}