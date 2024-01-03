package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Could not load envoirnment variables : ", err)
	}
	return os.Getenv(key)
}
