package helpers

import (
	"errors"
	"html"
	"regexp"
	"strings"
)

func Santize(data string) string {
	data = html.EscapeString(strings.TrimSpace(data))
	return data
}

func IsValidatePhoneNumber(phoneNumber string) error {
	const phoneNumberRegex = `^\+[1-9]\d{1,14}$`
	matched, err := regexp.MatchString(phoneNumberRegex, phoneNumber)
	if matched && err == nil {
		return nil
	}
	return errors.New("invalid phone number")
}
