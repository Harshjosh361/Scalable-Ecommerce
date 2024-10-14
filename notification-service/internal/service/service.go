package service

import (
	"fmt"
	"log"
	"os"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

// load env variables
// func init() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// }

type NotificationService struct {
	twilioClient *twilio.RestClient
}

func NewNotificationService() *NotificationService {
	accountSid := os.Getenv("Account_SID")
	authToken := os.Getenv("Auth_Token")

	if accountSid == "" || authToken == "" {
		log.Fatal("Twilio credentials are missing")
	}

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	return &NotificationService{
		twilioClient: client,
	}
}

func (ns *NotificationService) SendSMS(to string, body string) error {
	phoneNumber := os.Getenv("PhoneNumber")
	params := &twilioApi.CreateMessageParams{}
	params.SetFrom(phoneNumber)
	params.SetBody(body)
	params.SetTo(to)

	_, err := ns.twilioClient.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("failed to send message", err)
		return err
	}
	return nil
}
