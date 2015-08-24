package mpower

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SuiteTestCheckoutInvoice struct {
	suite.Suite
	mpower          *MPower
	checkoutInvoice *CheckoutInvoice
}

func (s *SuiteTestCheckoutInvoice) SetupSuite() {
	store := NewStore("Awesome Store")
	setup := NewSetupFromEnv()

	s.mpower = NewMPower(setup, store, "test")
	s.checkoutInvoice = NewCheckoutInvoice(s.mpower)
}

func (s *SuiteTestCheckoutInvoice) TestCreate() {
	s.checkoutInvoice.AddItem("Yam Phone", 1, 50.00, 50.00, "Hello World")
	s.checkoutInvoice.SetDescription("Hello World")
	s.checkoutInvoice.SetTotalAmount(50.00)

	resBody, _, err := s.checkoutInvoice.Create()
	fmt.Println(resBody.Token)
	assert.Nil(s.T(), err, "error is nil")
	assert.NotEmpty(s.T(), resBody.Token, "token was created"+resBody.Token)

	res, _, err2 := s.checkoutInvoice.Confirm(resBody.Token)
	fmt.Println(res.Status)
	assert.Nil(s.T(), err2, "error is nil")
	assert.NotEmpty(s.T(), res.Status, "Status found")
	assert.Equal(s.T(), "pending", res.Status, "status is the same")
}

func TestCheckoutInvoiceRunSuite(t *testing.T) {
	checkoutTester := new(SuiteTestCheckoutInvoice)
	suite.Run(t, checkoutTester)
}
