package types

type GeolocationData struct {
	Latitude	float64 `url:"latitude" validate:"required,latitude"`
	Longitude	float64 `url:"longitude" validate:"required,longitude"`
}