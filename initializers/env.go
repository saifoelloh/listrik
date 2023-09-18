package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GoDotEnvVariable() {
	if os.Getenv("APP_ENV") == "development" {
		// load .env file
		err := godotenv.Load(".env")

		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}
}
