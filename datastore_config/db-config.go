package datastore_config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GoDotEnvVar(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("%s", err)
	}
	return os.Getenv(key)
}