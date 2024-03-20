package env

import (
	"os"
	"strconv"
)

func GetEnvWithDefault(key, defaultValue string) string {

	res, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	return res
}

func GetEnvIntWithDefault(key string, defaultValue int) int {

	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	res, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return res
}
