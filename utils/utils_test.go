package utils

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DroppedHard/weather-forecast-api/types"
	"github.com/stretchr/testify/assert"
)

func TestValidateGeolocationData_ValidParams(t *testing.T) {
	// Mock a request with valid query parameters
	params := &types.GeolocationData{}
	req := httptest.NewRequest(http.MethodGet, "/test?Latitude=51.5074&Longitude=-0.1278", nil)

	err := ValidateGeolocationData(req, params)

	// Assert that there is no error
	assert.NoError(t, err)
	assert.Equal(t, 51.5074, params.Latitude)
	assert.Equal(t, -0.1278, params.Longitude)
}

func TestValidateGeolocationData_LatitudeOutOfRange(t *testing.T) {
	// Mock a request with invalid query parameters
	params := &types.GeolocationData{}
	req1 := httptest.NewRequest(http.MethodGet, "/test?Latitude=200&Longitude=-20", nil)
	req2 := httptest.NewRequest(http.MethodGet, "/test?Latitude=-200&Longitude=-20", nil)

	err1 := ValidateGeolocationData(req1, params)
	err2 := ValidateGeolocationData(req2, params)

	// Assert that the error is returned because the lat/long are out of range
	assert.Error(t, err1)
	assert.Contains(t, err1.Error(), "invalid query parameters")
	assert.Error(t, err2)
	assert.Contains(t, err2.Error(), "invalid query parameters")
}
func TestValidateGeolocationData_LongitudeOutOfRange(t *testing.T) {
	// Mock a request with invalid query parameters
	params := &types.GeolocationData{}
	req1 := httptest.NewRequest(http.MethodGet, "/test?Latitude=20&Longitude=-200", nil)
	req2 := httptest.NewRequest(http.MethodGet, "/test?Latitude=20&Longitude=200", nil)

	err1 := ValidateGeolocationData(req1, params)
	err2 := ValidateGeolocationData(req2, params)

	// Assert that the error is returned because the lat/long are out of range
	assert.Error(t, err1)
	assert.Contains(t, err1.Error(), "invalid query parameters")
	assert.Error(t, err2)
	assert.Contains(t, err2.Error(), "invalid query parameters")
}

func TestWriteJSON(t *testing.T) {
	w := httptest.NewRecorder()
	data := map[string]string{"message": "hello world"}
	expectedBody := `{"message":"hello world"}`
	err := WriteJSON(w, http.StatusOK, data)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestWriteError(t *testing.T) {
	w := httptest.NewRecorder()
	expectedBody := `{"error":"something went wrong"}`
	err := errors.New("something went wrong")

	WriteError(w, http.StatusInternalServerError, err)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.JSONEq(t, expectedBody, w.Body.String())
}
