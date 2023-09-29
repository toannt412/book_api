package service

import (
	"errors"
	"regexp"
	"testing"

	"github.com/stretchr/testify/suite"
)

// type GetAdminTestSuite struct {
// 	suite.Suite
// }

// func WrongID(ctx context.Context, id string) (*serialize.Admin, error) {
// 	// id = "650d4080be6273000b09ffea"
// 	// user, err := admin.GetAdminByID(ctx, id)
// 	return nil, nil
// }

const (
	IDRegex = "^[0-9a-fA-F]{24}$"
)
type ValidatorIDTestSuite struct {
	suite.Suite
}

func ValidateID(adminID string) error {
	matched, err := regexp.MatchString(IDRegex, adminID)
	if matched && err == nil {
		return nil
	}
	return errors.ErrUnsupported

}

func (suite *ValidatorIDTestSuite) TestInvalidLengthTooShortID() {
	err := ValidateID("fsfsdfd")
	suite.Equal(errors.ErrUnsupported, err)
}

func (suite *ValidatorIDTestSuite) TestInvalidLengthTooLongID() {
	err := ValidateID("auhkvaefvbiljfbvliefn1f151ew151ergw")
	suite.Equal(errors.ErrUnsupported, err)
}

func (suite *ValidatorIDTestSuite) TestInvalidContainsSpecialCharactersID() {
	err := ValidateID("/*-d*cdccdfv4fv156e=-9-]")
	suite.Equal(errors.ErrUnsupported, err)
}
func (suite *ValidatorIDTestSuite) TestEmtyID() {
	err := ValidateID("")
	suite.Equal(errors.ErrUnsupported, err)
}

func (suite *ValidatorIDTestSuite) TestValidID() {
	err := ValidateID("650d4080be6273000b09ffea")
	suite.Nil(err)
}

func TestValidatorIDTestSuite(t *testing.T) {
	suite.Run(t, new(ValidatorIDTestSuite))
}
