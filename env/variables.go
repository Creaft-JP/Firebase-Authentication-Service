package env

import (
	"os"
)

var ProjectId string

func init() {
	ProjectId = getOrElse("PROJECT_ID", "")
}

func getOrElse(key string, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
