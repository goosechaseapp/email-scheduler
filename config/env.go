package config

import "os"

// Returns the value of the environment variable with the given key.
// If the environment variable is not set, the default value is returned.
func Env(key string, defaultValue ...string) string {
	value, exists := os.LookupEnv(key)

	if !exists {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}

		return ""
	}

	return value
}
