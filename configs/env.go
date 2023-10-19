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
	FromPhone  string
	ToPhone    string
	ServiceSID string
	APIToken   string
	FromEmail  string
	ToEmail    string
	SMTPHost   string
	SMTPPORT   string
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
		FromPhone:  os.Getenv("FROM_PHONE"),
		ToPhone:    os.Getenv("TO_PHONE"),
		ServiceSID: os.Getenv("TWILIO_SERVICE_SID"),
		APIToken:   os.Getenv("API_TOKEN"),
		FromEmail:  os.Getenv("FROM_EMAIL"),
		ToEmail:    os.Getenv("TO_EMAIL"),
		SMTPHost:   os.Getenv("SMTP_HOST"),
		SMTPPORT:   os.Getenv("SMTP_PORT"),
	}
}
