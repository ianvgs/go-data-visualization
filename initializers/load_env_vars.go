package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	varMode := os.Getenv("GO_ENV")
	log.Println("MODE:", varMode)

	if os.Getenv("GO_ENV") != "production" {
		err := godotenv.Load()

		if err != nil {
			log.Fatal("Error loading dotenv file")
		}

	}

}
