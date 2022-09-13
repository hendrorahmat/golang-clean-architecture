package utils

import "os"

func GetEnv(name string) string {
	return os.Getenv(name)
}

func GetEnvWithDefaultValue(name string, defaultValue string) string {
	if os.Getenv(name) == "" {
		return defaultValue
	}
	return os.Getenv(name)
}
