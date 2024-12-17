package weather

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

	weatherResp, err := getWeeklyWeatherData(geoData, setForecastRequestParams)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	defer weatherResp.Body.Close()

	response, err := processWeatherData(weatherResp.Body)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	w.Header().Set("Content-Type", "application/json")
	utils.WriteJSON(w, http.StatusOK, response)
}

func (h *Handler) handleSummary(w http.ResponseWriter, r *http.Request) {
	var geoData types.GeolocationData

	if err := utils.ValidateGeolocationData(r, &geoData); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	weatherResp, err := getWeeklyWeatherData(geoData, setSummaryRequestParams)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	defer weatherResp.Body.Close()

	response, err := processSummaryData(weatherResp.Body)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	w.Header().Set("Content-Type", "application/json")
	utils.WriteJSON(w, http.StatusOK, response)
}

func getWeeklyWeatherData(geoData types.GeolocationData, setRequestParams types.SetParamsFunc) (*http.Response, error) {
	weatherUrl, err := url.Parse(config.Envs.WEATHER_API + "forecast")
	if err != nil {
		return nil, fmt.Errorf("failed to parse weather API URL: %w", err)
	}
	params := setRequestParams(geoData)
	weatherUrl.RawQuery = params.Encode()

	weatherReq, err := http.NewRequest("GET", weatherUrl.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create weather API request: %w", err)
	}
	weatherReq.Header.Add("Accept", "application/json")

	weatherResp, err := http.DefaultClient.Do(weatherReq)
	
	if err != nil {
		return nil, fmt.Errorf("failed to execute weather API request: %w", err)
	}
	if weatherResp.StatusCode != http.StatusOK {
		defer weatherResp.Body.Close()
		body, _ := io.ReadAll(weatherResp.Body)
		return nil, fmt.Errorf("weather API error: (%v): %s", weatherResp.StatusCode, body)
	}
	return weatherResp, nil
}
