package mpower

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type InvoiceSuiteTest struct {
	suite.Suite
	invoice Invoice
}

func (s *InvoiceSuiteTest) SetupSuite() {
	err, newStore := NewStore(map[string]string{
		"name":          "Awesome Store",
		"tagline":       "Easy shopping",
		"phoneNumber":   "0272271893",
		"postalAddress": "P.0. Box MP555, Accra",
		"logoURL":       "http://www.awesomestore.com.gh/logo.png",
	})

	assert.NoError(s.T(), err, "No Error")
	newSetup := NewSetup(map[string]string{
		"masterKey":  "34545-54565763-2323246-5455",
		"privateKey": "test_private_afdipfhpisfjroejr",
		"publicKey":  "test_public_fsofoufyrsfo",
		"token":      "dapifu09ur0jvsij",
		"mode":       "test",
	})

	s.invoice = Invoice{Setup: newSetup, Store: *newStore}
}

func (s *InvoiceSuiteTest) SetupTest() {
	s.invoice.AddItem("Bayere phone", 1, 20.00, 20.00, "It is a yam phone")
	s.invoice.AddTax("VAT", 30.00)
}

func (s *InvoiceSuiteTest) TearDownTest() {
	s.invoice.RemoveItem("Bayere phone")
	s.invoice.RemoveTax("VAT")
}

func (s *InvoiceSuiteTest) TestAddItem() {
	assert.Equal(s.T(), 1, len(s.invoice.InvoiceIn.ItemsArr), "items length is 1")
	assert.NotNil(s.T(), s.invoice.InvoiceIn.ItemsArr, "invoice items contains an item")
	assert.Equal(s.T(), "Bayere phone", s.invoice.InvoiceIn.ItemsArr[0].Name, "Bayere phone is there")
	assert.Equal(s.T(), 1, s.invoice.InvoiceIn.ItemsArr[0].Quantity, "Bayere phone is only 1")
	s.invoice.AddItem("Sobolo", 1, 5.00, 5.00, "The yoghurt of Africa")
	assert.Equal(s.T(), 2, len(s.invoice.InvoiceIn.ItemsArr), "items length is 1")
}

func (s *InvoiceSuiteTest) TestRemoveItem() {
	s.invoice.RemoveItem("Sobolo")
	assert.Equal(s.T(), 1, len(s.invoice.InvoiceIn.ItemsArr), "items length is 1")
}

func (s *InvoiceSuiteTest) TestClearAllItems() {
	assert.Equal(s.T(), 1, len(s.invoice.InvoiceIn.ItemsArr), "items length is 1")
	assert.NotNil(s.T(), s.invoice.InvoiceIn.ItemsArr, "invoice items contains an item")
	assert.Equal(s.T(), "Bayere phone", s.invoice.InvoiceIn.ItemsArr[0].Name, "Bayere phone is there")
	assert.Equal(s.T(), 1, s.invoice.InvoiceIn.ItemsArr[0].Quantity, "Bayere phone is only 1")

	s.invoice.ClearAllItems()
	assert.Equal(s.T(), 0, len(s.invoice.InvoiceIn.ItemsArr), "no item exists")
}

func (s *InvoiceSuiteTest) TestAddTax() {
	assert.Equal(s.T(), 1, len(s.invoice.InvoiceIn.TaxesArr), "taxeslength is 1")
	assert.NotNil(s.T(), s.invoice.InvoiceIn.TaxesArr, "invoice taxes contains an tax")
	assert.Equal(s.T(), "VAT", s.invoice.InvoiceIn.TaxesArr[0].Name, "Tax is there")
	assert.Equal(s.T(), float32(30.00), s.invoice.InvoiceIn.TaxesArr[0].Amount, "Tax is there")
	s.invoice.AddTax("NHIL", 500.00)
	assert.Equal(s.T(), 2, len(s.invoice.InvoiceIn.TaxesArr), "taxes length is 2")
}

func (s *InvoiceSuiteTest) TestRemoveTax() {
	s.invoice.RemoveTax("NHIL")
	assert.Equal(s.T(), 1, len(s.invoice.InvoiceIn.TaxesArr), "taxes length is 1")
}

func (s *InvoiceSuiteTest) TestClear() {
	s.invoice.Clear()
	assert.Equal(s.T(), 0, len(s.invoice.InvoiceIn.TaxesArr), "taxes length is 0")
	assert.Equal(s.T(), 0, len(s.invoice.InvoiceIn.ItemsArr), "items length is 0")
}

func (s *InvoiceSuiteTest) TestSetDescription() {
	s.invoice.SetDescription("desc")
	assert.Equal(s.T(), "desc", s.invoice.InvoiceIn.Description, "description is equal")
}

func (s *InvoiceSuiteTest) TestSetTotalAmount() {
	s.invoice.SetTotalAmount(50.00)
	assert.Equal(s.T(), float32(50.00), s.invoice.InvoiceIn.TotalAmount, "total amount is 50.00")
}

func (s *InvoiceSuiteTest) TestCustomData() {
	s.invoice.SetCustomData("me", "yeah")
	assert.Equal(s.T(), "yeah", s.invoice.CustomData["me"], "me is yeah")
}

func (s *InvoiceSuiteTest) TestPrepForRequest() {
	s.invoice.PrepareForRequest()
	assert.Equal(s.T(), "VAT", s.invoice.InvoiceIn.Taxes["tax_0"].Name, "Tax name is VAT")
	assert.Equal(s.T(), "Bayere phone", s.invoice.InvoiceIn.Items["item_0"].Name, "Item name is bayere")
}

func TestInvoiceSuiteTest(t *testing.T) {
	invoiceSuiteTester := new(InvoiceSuiteTest)
	suite.Run(t, invoiceSuiteTester)
}
