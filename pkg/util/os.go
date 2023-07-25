package util

import "os"

// GetEnv returns the value of an environment variable or panics
// if it is not set.
func GetEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		panic("env var " + key + " is not set")
	}
	return value
}

// GetEnvFallback returns the value of an environment variable or a fallback.
func GetEnvFallback(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return value
}
