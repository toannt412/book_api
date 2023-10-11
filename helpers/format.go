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

func ValidatePhoneNumber(phoneNumber string) error {
	const phoneNumberRegex = "^[0-9]{4,13}$"
	matched, err := regexp.MatchString(phoneNumberRegex, phoneNumber)
	if matched && err == nil {
		return nil
	}
	return errors.New("invalid phone number")
}
