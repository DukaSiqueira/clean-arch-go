package environments

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error landing .env file")
	}

	log.Println("Successfully loaded .env file")
}
