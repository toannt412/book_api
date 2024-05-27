package test

import (
	"errors"
	"regexp"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	EmailRegex = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
)

type ValidatorEmailTestSuite struct {
	suite.Suite
}

func ValidateEmail(email string) error {
	matched, err := regexp.MatchString(EmailRegex, email)
	if matched && err == nil {
		return nil
	}
	return errors.ErrUnsupported
}

func (s *ValidatorEmailTestSuite) TestEmtyEmail() {
	err := ValidateEmail("")
	s.Equal(errors.ErrUnsupported, err)
}

func (s *ValidatorEmailTestSuite) TestInvalidEmailWithoutAt() {
	err := ValidateEmail("testemail.com")
	s.Equal(errors.ErrUnsupported, err)
}

func (s *ValidatorEmailTestSuite) TestInvalidEmailWithoutUsername() {
	err := ValidateEmail("@gmail.com")
	s.Equal(errors.ErrUnsupported, err)
}

func (s *ValidatorEmailTestSuite) TestInvalidEmailWithoutDomain() {
	err := ValidateEmail("test@.com")
	s.Equal(errors.ErrUnsupported, err)
}

func (s *ValidatorEmailTestSuite) TestValidEmail() {
	err := ValidateEmail("test@gmail.com")
	s.Nil(err)
}

func TestValidatorEmailTestSuite(t *testing.T) {
	suite.Run(t, new(ValidatorEmailTestSuite))
}
