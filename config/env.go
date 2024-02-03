package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	ENV_FILE = ".env"
	PORT     = "PORT"
)

func init() {
	if err := godotenv.Load(ENV_FILE); err != nil {
		log.Fatal(err)
	}
}

func Get(key string) string {
	return os.Getenv(key)
}
