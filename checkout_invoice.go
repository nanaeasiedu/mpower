package mpower

import (
	"github.com/jmcvetta/napping"
)

// CheckoutInvoice holds all the data related to checkout invoice
// Invoice is an embedded struct, so all methods of Invoice can be called on it
type CheckoutInvoice struct {
	Invoice
	baseURL string
	mpower  *MPower
}

type CheckoutInvoiceRequest struct {
	Invoice
}

// CheckoutInvoiceResponse is the response data as specified by the mpower
// It retrieves the response json data and stores it on the checkout invoice object
type CheckoutInvoiceResponse struct {
	Response
	Token string `json:"token,omitempty"`
}

// CheckoutInvoiceStatus holds all the data related to status of an invoice created on mpower
type CheckoutInvoiceStatus struct {
	Response
	Status string `json:"status,omitempty"`
}

// Create - creates a new invoice on mpower
func (c *CheckoutInvoice) Create() (*CheckoutInvoiceResponse, *napping.Response, error) {
	payload := &CheckoutInvoiceRequest{Invoice: c.Invoice}
	responseBody := &CheckoutInvoiceResponse{}

	response, err := c.mpower.NewRequest("POST", c.baseURL+"/create", payload, responseBody, nil)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

// Confirm - This confirms the token status
func (c *CheckoutInvoice) Confirm(token string) (string, error) {
	responseBody := &CheckoutInvoiceStatus{}
	response, err := c.mpower.NewRequest("GET", c.baseURL+"/confirm", nil, responseBody, nil)

	if err != nil {
		return "", err
	}

	return response.RawText(), nil
}

// NewCheckoutInvoice - create a new checkout instance
//
// Example.
//     checkout := mpower.NewCheckoutInvoice(mpower)
func NewCheckoutInvoice(mp *MPower) *CheckoutInvoice {
	checkoutInvoice := &CheckoutInvoice{Invoice: Invoice{Setup: mp.setup, Store: *mp.store}}
	checkoutInvoice.mpower = mp
	checkoutInvoice.baseURL = mp.baseURL + "/checkout-invoice"
	return checkoutInvoice
}
