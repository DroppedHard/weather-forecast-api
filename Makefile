build:
	@go build -o bin/weather-forecast-api cmd/main.go

test: 
	@go test -v ./...

run: build
	@./bin/weather-forecast-api