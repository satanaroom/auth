package env

import (
	"os"
	"strconv"
)

// ToInt converts env variable to int
func ToInt(value *int, key string, defaultValue int) {
	*value = defaultValue

	envValue, exists := os.LookupEnv(key)
	if !exists || envValue == "" {
		return
	}

	if res, err := strconv.Atoi(envValue); err == nil {
		*value = res
	}
}

// ToFloat converts env variable to float
func ToFloat(value *float64, key string, defaultValue float64) {
	*value = defaultValue

	envValue, exists := os.LookupEnv(key)
	if !exists || envValue == "" {
		return
	}

	if res, err := strconv.ParseFloat(envValue, 64); err == nil {
		*value = res
	}
}

// ToBool converts env variable to bool
func ToBool(value *bool, key string, defaultValue bool) {
	*value = defaultValue

	envValue, exists := os.LookupEnv(key)
	if !exists || envValue == "" {
		return
	}

	if res, err := strconv.ParseBool(envValue); err == nil {
		*value = res
	}
}

// ToString converts env variable to string
func ToString(value *string, key string, defaultValue string) {
	*value = getEnv(key, defaultValue)
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists || value != "" {
		return value
	}

	return defaultValue
}
