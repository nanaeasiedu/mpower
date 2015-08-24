package mpower

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SuiteTestDirectPay struct {
	suite.Suite
	mpower    *MPower
	directPay *DirectPay
}

func (s *SuiteTestDirectPay) SetupSuite() {
	store := NewStore("Awesome Store")
	setup := NewSetupFromEnv()

	s.mpower = NewMPower(setup, store, "test")
	s.directPay = NewDirectPay(s.mpower)
}

func (s *SuiteTestDirectPay) TestDirectPay() {
	resBody, _, err := s.directPay.CreditAccount("Ngenerio", 100.00)
	fmt.Println(resBody)
	fmt.Println(resBody.TransactionID)
	assert.Nil(s.T(), err, "error is nil")
	assert.NotEmpty(s.T(), resBody.TransactionID, "token was created"+resBody.TransactionID)
}

func TestDirectPayRunSuite(t *testing.T) {
	checkoutTester := new(SuiteTestDirectPay)
	suite.Run(t, checkoutTester)
}
