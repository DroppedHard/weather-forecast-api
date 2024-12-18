package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	os.Setenv("TEST_KEY", "test_value")
	defer os.Unsetenv("TEST_KEY")

	result1 := getEnv("TEST_KEY", "default_value")
	result2 := getEnv("NON_EXISTENT_KEY", "default_value")

	assert.Equal(t, "test_value", result1, "Expected value from environment variable")
	assert.Equal(t, "default_value", result2, "Expected default value")
}

func TestInitConfig_WithEnv(t *testing.T) {
	os.Setenv("PORT", "9090")
	os.Setenv("WEATHER_API", "https://api.custom.com/")
	os.Setenv("ORIGINS_ALLOWED", "https://example.com")
	defer os.Unsetenv("PORT")
	defer os.Unsetenv("WEATHER_API")
	defer os.Unsetenv("ORIGINS_ALLOWED")

	Envs = initConfig()

	assert.Equal(t, "9090", Envs.PORT, "Expected PORT to be set to 9090")
	assert.Equal(t, "https://api.custom.com/", Envs.WEATHER_API, "Expected WEATHER_API to be set to custom URL")
	assert.Equal(t, "https://example.com", Envs.ORIGINS_ALLOWED, "Expected ORIGINS_ALLOWED to be set to https://example.com")
}

func TestInitConfig_WithFallback(t *testing.T) {
	os.Unsetenv("PORT")
	os.Unsetenv("WEATHER_API")
	os.Unsetenv("ORIGINS_ALLOWED")

	Envs = initConfig()

	assert.Equal(t, "8080", Envs.PORT, "Expected PORT to fallback to 8080")
	assert.Equal(t, "https://api.open-meteo.com/v1/", Envs.WEATHER_API, "Expected WEATHER_API to fallback to default URL")
	assert.Equal(t, "scheme://dns[:port]", Envs.ORIGINS_ALLOWED, "Expected ORIGINS_ALLOWED to fallback to default value")
}
