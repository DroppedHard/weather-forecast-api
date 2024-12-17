package weather

import (
	"net/http"

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

	response, err := ProcessWeatherData(weatherResp.Body)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	w.Header().Set("Content-Type", "application/json")
	utils.WriteJSON(w, http.StatusOK, response)
}

func (h *Handler) handleSummary(w http.ResponseWriter, r *http.Request) {
	// validate parameters from request

}
