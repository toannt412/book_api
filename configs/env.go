package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Config ConfigSchema

type ConfigSchema struct {
	MongoURI   string
	Port       string
	AccountSID string
	AuthToken  string
}

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Config = ConfigSchema{
		MongoURI:   os.Getenv("MONGOURI"),
		Port:       os.Getenv("PORT"),
		AccountSID: os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken:  os.Getenv("TWILIO_AUTH_TOKEN"),
	}
}
