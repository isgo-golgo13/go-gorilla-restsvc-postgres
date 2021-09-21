package db_config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GoDotEnvVar(key string) string {
	err := godotenv.Load("./db_config/.env")
	if err != nil {
		log.Fatalf("%s", err)
	}
	return os.Getenv(key)
}