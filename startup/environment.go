package startup

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvironment() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}
