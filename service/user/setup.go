package user

import (
	third_party "bookstore/3P"
	"bookstore/dao/user"
)

type UserService struct {
	userRepo *user.UserRepository
	twilio   *third_party.Twilio
	brevo    *third_party.Brevo
	aws      *third_party.AWSService
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: user.NewUserRepository(),
		twilio:   third_party.NewTwilio(),
		brevo:    third_party.NewBrevo(),
		aws:      third_party.NewAWSService(),
	}
}
