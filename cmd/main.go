package main

import (
	"log"

	"github.com/DroppedHard/weather-forecast-api/cmd/api"
	"github.com/DroppedHard/weather-forecast-api/config"
)

func main() {
	server := api.NewAPIServer(":" + config.Envs.PORT, nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}