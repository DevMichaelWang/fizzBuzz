package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func GoDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		return "env_err"
	}

	return os.Getenv(key)
}
