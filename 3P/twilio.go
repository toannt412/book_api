package third_party

import (
	"bookstore/configs"
	"encoding/json"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type Twilio struct {
	client *twilio.RestClient
}

func NewTwilio() *Twilio {
	return &Twilio{
		client: twilio.NewRestClientWithParams(twilio.ClientParams{
			Username: configs.Config.AccountSID,
			Password: configs.Config.AuthToken,
		}),
	}
}

func (t *Twilio) SendOTPByPhone(otp string) (string, error) {
	params := &twilioApi.CreateMessageParams{}
	params.SetTo(configs.Config.ToPhone)
	params.SetFrom(configs.Config.FromPhone)
	params.SetBody("Your verification code is: " + otp)

	resp, err := t.client.Api.CreateMessage(params)
	if err != nil {
		return "", err
	} else {
		response, _ := json.Marshal(*resp)
		return string(response), nil
	}
}
