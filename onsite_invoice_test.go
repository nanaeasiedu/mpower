package mpower

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SuiteTestOnsiteInvoice struct {
	suite.Suite
	mpower        *MPower
	onsiteInvoice *OnsiteInvoice
}

func (s *SuiteTestOnsiteInvoice) SetupSuite() {
	store := NewStore("Awesome Store")
	setup := NewSetupFromEnv()

	s.mpower = NewMPower(setup, store, "test")
	s.onsiteInvoice = NewOnsiteInvoice(s.mpower)
}

func (s *SuiteTestOnsiteInvoice) TestOnsiteCreate() {
	s.onsiteInvoice.AddItem("Yam Phone", 1, 50.00, 50.00, "Hello World")
	s.onsiteInvoice.SetDescription("Hello World")
	s.onsiteInvoice.SetTotalAmount(50.00)

	resBody, _, err := s.onsiteInvoice.Create("Ngenerio")
	fmt.Println(resBody)
	fmt.Println(resBody.InvoiceToken)
	assert.Nil(s.T(), err, "error is nil")
	assert.NotEmpty(s.T(), resBody.Token, "token was created"+resBody.Token)
	assert.NotEmpty(s.T(), resBody.InvoiceToken, "token was created"+resBody.InvoiceToken)
}

func TestOnsiteInvoiceRunSuite(t *testing.T) {
	checkoutTester := new(SuiteTestOnsiteInvoice)
	suite.Run(t, checkoutTester)
}
