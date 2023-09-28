package controllers

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	//r.GET("/admin/:adminId", GetAAdmin())
	return r
}

func TestGetAAdmin(t *testing.T) {
	router := setupRouter()
	
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/admin/650d4080be6273000b09ffea", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "success", w.Body.String())
}

// const (
// 	phoneNumberRegex  = "^[0-9]{4,13}$"
//   )

//   type ValidatorPhoneNumberTestSuite struct {
// 	suite.Suite
//   }

//   func ValidatePhoneNumber(phoneNumber string) error {
// 	matched, err := regexp.MatchString(phoneNumberRegex, phoneNumber)
// 	if matched && err == nil {
// 	  return nil
// 	}
// 	return errors.ErrValidatePhone
//   }

//   func (suite *ValidatorPhoneNumberTestSuite) TestInvalidLengthPhoneNumber() {
// 	err := ValidatePhoneNumber("123")
// 	suite.Equal(errors.ErrValidatePhone, err)
//   }

//   func (suite *ValidatorPhoneNumberTestSuite) TestInvalidLengthPhoneNumberCharacter() {
// 	err := ValidatePhoneNumber("fsfsdfd")
// 	suite.Equal(errors.ErrValidatePhone, err)
//   }

//   func (suite *ValidatorPhoneNumberTestSuite) TestValidPhoneNumber() {
// 	err := ValidatePhoneNumber("0958545212")
// 	suite.Nil(err)
//   }

//   func TestValidatorPhoneNumberTestSuite(t *testing.T) {
// 	suite.Run(t, new(ValidatorPhoneNumberTestSuite))
//   }
