package utils

import (
	"log"
	"os"
)

func GetEnvVariables(key string) string {
	value := os.Getenv(key)

	if value == "" {
		log.Fatalf("warning: Environment variable %s not provided", key)
		return ""
	}

	return value

}
