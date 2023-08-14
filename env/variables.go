package env

import (
	"github.com/joho/godotenv"
	"os"
)

var ProjectId string

func init() {
	_ = godotenv.Load()
	ProjectId = getOrElse("PROJECT_ID", "")
}

func getOrElse(key string, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
