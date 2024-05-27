package third_party

import (
	"bookstore/configs"
	"strconv"

	"gopkg.in/gomail.v2"
)

type Brevo struct {
	mail *gomail.Dialer
}

func NewBrevo() *Brevo {
	port, _ := strconv.Atoi(configs.Config.SMTPPORT)
	return &Brevo{
		mail: gomail.NewDialer(configs.Config.SMTPHost, port, configs.Config.FromEmail, configs.Config.APIToken),
	}
}
func (m *Brevo) SendOTPByEmail(otp string, email string) (string, error) {
	from := configs.Config.FromEmail

	msg := gomail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", email)
	msg.SetHeader("Subject", "OTP for forgot password")
	// text/html for a html email
	msg.SetBody("text/plain", "Your OTP is: "+otp)

	//m.mail.DialAndSend(msg)

	// Send the email
	if err := m.mail.DialAndSend(msg); err != nil {
		panic(err)
	}
	return "send email success", nil
}
