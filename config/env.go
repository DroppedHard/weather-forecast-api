package config

import (
	"log"
	"os"

	"github.com/lpernett/godotenv"
)

type Config struct {
	PORT		string
	WEATHER_API	string
}

var Envs = initConfig()

func initConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	return Config{
		PORT: getEnv("PORT", "8080"),
		WEATHER_API: getEnv("WEATHER_API", "https://api.open-meteo.com/v1/"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}