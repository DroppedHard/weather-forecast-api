package user

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/DroppedHard/weather-forecast-api/config"
	"github.com/DroppedHard/weather-forecast-api/types"
	"github.com/DroppedHard/weather-forecast-api/utils"
	"github.com/gorilla/mux"
)
type Handler struct {

}

func NewHandler() *Handler {
	return &Handler{}
}
func (h *Handler) RegisterRoutes (router *mux.Router) {
	router.HandleFunc("/forecast", h.handleForecast).Methods("GET")
	router.HandleFunc("/summary", h.handleSummary).Methods("GET")
}

func (h *Handler) handleForecast(w http.ResponseWriter, r *http.Request) {
	var geoData types.GeolocationData

	if err := utils.ValidateGeolocationData(r, &geoData); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	weatherResp, err := GetWeatherForecast(geoData)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	defer weatherResp.Body.Close()

	body, err := io.ReadAll(weatherResp.Body)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to read weather API response: %w", err))
		return
	}

	// Print the response body
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
func (h *Handler) handleSummary(w http.ResponseWriter, r *http.Request) {
	// validate parameters from request

}

func GetWeatherForecast(geoData types.GeolocationData) (*http.Response, error) {
	// prepare wheather forecast URL and params
	weatherUrl, err := url.Parse(config.Envs.WeatherApi + "forecast")
	if err != nil {
		return nil, fmt.Errorf("failed to parse weather API URL: %w", err)
	}
	params := utils.SetRequestParams(geoData)
	weatherUrl.RawQuery = params.Encode()

	// Create the request, and send it
	weatherReq, err := http.NewRequest("GET", weatherUrl.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create weather API request: %w", err)
	}
	weatherReq.Header.Add("Accept", "application/json")

	weatherResp, err := http.DefaultClient.Do(weatherReq)
	
	// verify whether everything went properly
	if err != nil {
		return nil, fmt.Errorf("failed to execute weather API request: %w", err)
	}
	if weatherResp.StatusCode != http.StatusOK {
		weatherResp.Body.Close()
		return nil, fmt.Errorf("weather API returned unexpected status: %d", weatherResp.StatusCode)
	}
	return weatherResp, nil
}