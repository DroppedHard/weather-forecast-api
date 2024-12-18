package utils

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"

	"github.com/DroppedHard/weather-forecast-api/types"
	"github.com/go-playground/form"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func init() {
	errLat := validate.RegisterValidation("latitude",func(f1 validator.FieldLevel) bool {
		lat, ok := f1.Field().Interface().(float64)
		return ok && lat >= -90 && lat <= 90
	})
	errLong := validate.RegisterValidation("longitude", func(fl validator.FieldLevel) bool {
		lon, ok := fl.Field().Interface().(float64)
		return ok && lon >= -180 && lon <= 180
	})
	if errLat != nil || errLong != nil {
		fmt.Println("failed to register validation - latitude: ", errLat," longitude: ", errLong)
		return 
	}
}


func ValidateGeolocationData(r *http.Request, params *types.GeolocationData) error {
	query := r.URL.Query()

	if err := form.NewDecoder().Decode(params, query); err != nil {
		return fmt.Errorf("failed to decode query parameters: %s", err)
	}

	if err := validate.Struct(params); err != nil {
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

func RoundFloat(val float64, precision uint) float64 {
    ratio := math.Pow(10, float64(precision))
    return math.Round(val*ratio) / ratio
}




