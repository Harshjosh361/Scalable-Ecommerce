package service

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

// load env variables
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

type NotificationService struct {
	twilioClient *twilio.RestClient
}

func NewNotificationService() *NotificationService {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: os.Getenv("Account_SID"),
		Password: os.Getenv("Auth_Token"),
	})

	return &NotificationService{
		twilioClient: client,
	}
}

func (ns *NotificationService) SendSMS(to string, body string) error {
	params := &twilioApi.CreateMessageParams{}
	params.SetFrom(os.Getenv("PhoneNumber"))
	params.SetBody(body)
	params.SetTo(to)

	_, err := ns.twilioClient.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("failed to send message", err)
		return err
	}
	return nil
}
