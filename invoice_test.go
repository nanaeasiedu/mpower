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
	newStore := NewStore("Awesome Store")

	newSetup := NewSetup("43434-54545-45454-545432", "test_private_auhidaudvbirbyyrieoib", "test_public_iopjasdioppdadipjoasd", "ioapdojdifouw8h")

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
	assert.Equal(s.T(), 1, len(s.invoice.InvoiceData.ItemsArr), "items length is 1")
	assert.NotNil(s.T(), s.invoice.InvoiceData.ItemsArr, "invoice items contains an item")
	assert.Equal(s.T(), "Bayere phone", s.invoice.InvoiceData.ItemsArr[0].Name, "Bayere phone is there")
	assert.Equal(s.T(), 1, s.invoice.InvoiceData.ItemsArr[0].Quantity, "Bayere phone is only 1")
	s.invoice.AddItem("Sobolo", 1, 5.00, 5.00, "The yoghurt of Africa")
	assert.Equal(s.T(), 2, len(s.invoice.InvoiceData.ItemsArr), "items length is 1")
}

func (s *InvoiceSuiteTest) TestRemoveItem() {
	s.invoice.RemoveItem("Sobolo")
	assert.Equal(s.T(), 1, len(s.invoice.InvoiceData.ItemsArr), "items length is 1")
}

func (s *InvoiceSuiteTest) TestClearAllItems() {
	assert.Equal(s.T(), 1, len(s.invoice.InvoiceData.ItemsArr), "items length is 1")
	assert.NotNil(s.T(), s.invoice.InvoiceData.ItemsArr, "invoice items contains an item")
	assert.Equal(s.T(), "Bayere phone", s.invoice.InvoiceData.ItemsArr[0].Name, "Bayere phone is there")
	assert.Equal(s.T(), 1, s.invoice.InvoiceData.ItemsArr[0].Quantity, "Bayere phone is only 1")

	s.invoice.ClearAllItems()
	assert.Equal(s.T(), 0, len(s.invoice.InvoiceData.ItemsArr), "no item exists")
}

func (s *InvoiceSuiteTest) TestAddTax() {
	assert.Equal(s.T(), 1, len(s.invoice.InvoiceData.TaxesArr), "taxeslength is 1")
	assert.NotNil(s.T(), s.invoice.InvoiceData.TaxesArr, "invoice taxes contains an tax")
	assert.Equal(s.T(), "VAT", s.invoice.InvoiceData.TaxesArr[0].Name, "Tax is there")
	assert.Equal(s.T(), float32(30.00), s.invoice.InvoiceData.TaxesArr[0].Amount, "Tax is there")
	s.invoice.AddTax("NHIL", 500.00)
	assert.Equal(s.T(), 2, len(s.invoice.InvoiceData.TaxesArr), "taxes length is 2")
}

func (s *InvoiceSuiteTest) TestRemoveTax() {
	s.invoice.RemoveTax("NHIL")
	assert.Equal(s.T(), 1, len(s.invoice.InvoiceData.TaxesArr), "taxes length is 1")
}

func (s *InvoiceSuiteTest) TestClear() {
	s.invoice.Clear()
	assert.Equal(s.T(), 0, len(s.invoice.InvoiceData.TaxesArr), "taxes length is 0")
	assert.Equal(s.T(), 0, len(s.invoice.InvoiceData.ItemsArr), "items length is 0")
}

func (s *InvoiceSuiteTest) TestSetDescription() {
	s.invoice.SetDescription("desc")
	assert.Equal(s.T(), "desc", s.invoice.InvoiceData.Description, "description is equal")
}

func (s *InvoiceSuiteTest) TestSetTotalAmount() {
	s.invoice.SetTotalAmount(50.00)
	assert.Equal(s.T(), float32(50.00), s.invoice.InvoiceData.TotalAmount, "total amount is 50.00")
}

func (s *InvoiceSuiteTest) TestCustomData() {
	s.invoice.SetCustomData("me", "yeah")
	assert.Equal(s.T(), "yeah", s.invoice.CustomData["me"], "me is yeah")
}

func (s *InvoiceSuiteTest) TestPrepForRequest() {
	s.invoice.PrepareForRequest()
	assert.Equal(s.T(), "VAT", s.invoice.InvoiceData.Taxes["tax_0"].Name, "Tax name is VAT")
	assert.Equal(s.T(), "Bayere phone", s.invoice.InvoiceData.Items["item_0"].Name, "Item name is bayere")
}

func TestInvoiceSuiteTest(t *testing.T) {
	invoiceSuiteTester := new(InvoiceSuiteTest)
	suite.Run(t, invoiceSuiteTester)
}
