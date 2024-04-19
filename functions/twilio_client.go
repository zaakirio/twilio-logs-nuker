package Handlers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
)

func InitClient() *twilio.RestClient {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")

	if accountSid == "" || authToken == "" {
		log.Fatal("TWILIO_ACCOUNT_SID and TWILIO_AUTH_TOKEN must be set in .env file")
	}

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
	return client
}
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
