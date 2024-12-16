package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/DroppedHard/weather-forecast-api/types"
	"github.com/go-playground/form"
	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func init() {
	errLat := Validate.RegisterValidation("latitude",func(f1 validator.FieldLevel) bool {
		lat, ok := f1.Field().Interface().(float64)
		return ok && lat >= -90 && lat <= 90
	})
	errLong := Validate.RegisterValidation("longitude", func(fl validator.FieldLevel) bool {
		lon, ok := fl.Field().Interface().(float64)
		return ok && lon >= -180 && lon <= 180
	})
	if errLat != nil || errLong != nil {
		fmt.Println("failed to register validation - latitude: ", errLat," longitude: ", errLong)
		return 
	}
}

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}


func ValidateGeolocationData(r *http.Request, params *types.GeolocationData) error {
	query := r.URL.Query()

	if err := form.NewDecoder().Decode(params, query); err != nil {
		return fmt.Errorf("failed to decode query parameters: %s", err)
	}

	if err := Validate.Struct(params); err != nil {
		return fmt.Errorf("invalid query parameters: %v", err)
	}

	return nil

}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	fmt.Println("Error: ", err)
	WriteJSON(w, status, map[string]string{"error":err.Error()})
}

var defaultParams = map[string]string{
	"daily":      "weather_code,temperature_2m_max,temperature_2m_min,daylight_duration",
	"timezone":   "GMT",
	"start_date": "2024-12-16",
	"end_date":   "2024-12-22",
}

func SetRequestParams(geoData types.GeolocationData) (params url.Values) {
	params = url.Values{}
	params.Add("latitude", strconv.FormatFloat(geoData.Latitude, 'f', -1, 64))
	params.Add("longitude", strconv.FormatFloat(geoData.Longitude, 'f', -1, 64))
	for key, value := range defaultParams {
		params.Add(key, value)
	}
	return
}
