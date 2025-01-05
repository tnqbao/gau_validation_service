package providers

import (
	"fmt"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	"math/rand"
	"os"
	"time"
)

func GenCaptchaCode() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var codes [6]byte
	for i := 0; i < 6; i++ {
		codes[i] = uint8(48 + r.Intn(10))
	}

	return string(codes[:])
}

func SendSMSWithTwilio(phone string, message string) (string, error) {
	var (
		TwilioAccountSID  = os.Getenv("TWILIO_ACCOUNT_SID")
		TwilioAuthToken   = os.Getenv("TWILIO_AUTH_TOKEN")
		TwilioPhoneNumber = os.Getenv("TWILIO_PHONE_NUMBER")
	)

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: TwilioAccountSID,
		Password: TwilioAuthToken,
	})

	params := &openapi.CreateMessageParams{}
	params.SetTo(phone)
	params.SetFrom(TwilioPhoneNumber)
	params.SetBody(message)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		return "", fmt.Errorf("failed to send SMS: %v", err)
	}

	return *resp.Sid, nil
}
