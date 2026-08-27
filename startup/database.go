package startup

import (
	"log"
	"os"
)

func InitializeDatabase() {
	dbURL := os.Getenv("MONGODB_URI")
	log.Println("DB:", dbURL)
}
