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
		log.Fatal("Error loading .env file")
	}
	
	return Config{
		PORT: getEnv("PORT", "8080"),
		WEATHER_API: getEnv("WEATHER_API", "https://api.open-meteo.com/v1/"),
		ORIGINS_ALLOWED: getEnv("ORIGINS_ALLOWED", "scheme://dns[:port]"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}