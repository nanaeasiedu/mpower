package mpower

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
)

// CheckoutInvoice holds all the data related to checkout invoice
// Invoice is an embedded struct, so all methods of Invoice can be called on it
type CheckoutInvoice struct {
	Invoice
	baseURL      string `json:"-"`
	ResponseCode string `json:"-"`
	ResponseText string `json:"-"`
	Description  string `json:"-"`
	Token        string `json:"-"`
	Status       string `json:"-"`
}

// The response data as specified by the mpower
// It retrieves the response json data and stores it on the checkout invoice object
type responseJSONCheckout struct {
	ResponseCode string `json:"response_code"`
	ResponseText string `json:"response_text"`
	Description  string `json:"description"`
	Token        string `json:"token"`
}

// stat holds all the data related to status of an invoice created on mpower
type stat struct {
	status       string `json:"status"`
	responseCode string `json:"response_code"`
}

// Create - creates a new invoice on mpower
// Returns `boolean` and `error`
// The `boolean` is used to determine if an error was encountered while making the request
//
// Example.
//    if ok, err := checkout.Create(); ok {
//      //do something with the response info on the checkout instance
//      fmt.Printf("%s %s %s %s\n\n", checkout.ResponseCode, checkout.ResponseText, checkout.Description, checkout.Token)
//    } else {
//      //there was an error
//    }
func (c *CheckoutInvoice) Create() (bool, error) {
	var respJSON responseJSONCheckout
	req := gorequest.New()

	c.PrepareForRequest()
	req.Post(c.baseURL + "/create")

	for key, val := range c.Setup.Headers {
		req.Set(key, val)
	}

	if content, err := json.Marshal(c.Invoice); err != nil {
		return false, err
	} else {
		req.Send(bytes.NewBuffer(content).String())
	}

	if resp, body, err := req.End(); err != nil {
		fmt.Errorf("%v", err)
		c.Status = resp.Status
		return false, fmt.Errorf("%v", err)
	} else {
		if err := json.Unmarshal(bytes.NewBufferString(body).Bytes(), &respJSON); err != nil {
			return false, err
		}

		if respJSON.ResponseCode == "00" {
			c.ResponseText = respJSON.ResponseText
			c.ResponseCode = respJSON.ResponseCode
			c.Description = respJSON.Description
			c.Token = respJSON.Token

			return true, nil
		}

		return false, fmt.Errorf("Failed to create invoice with error : %s", body)
	}
}

// GetInvoiceURL - get the invoice's url from the response
//
// Example.
//    str := checkout.GetInvoiceURL()
func (c *CheckoutInvoice) GetInvoiceURL() string {
	if c.Token == "" {
		panic("Token currently not available")
	}

	return c.ResponseText
}

// Confirm - This confirms the token status
//
// Example.
//     str, err := checkout.Confirm("434-5455-adf4-fgt5")
func (c *CheckoutInvoice) Confirm(token string) (string, error) {
	var status stat
	req := gorequest.New()

	req.Get(c.baseUrl + "/confirm/" + token)
	for key, val := range c.Setup.Headers {
		req.Set(key, val)
	}

	if resp, body, err := req.End(); err != nil {
		c.Status = resp.Status
		return "", fmt.Errorf("%v", err)
	} else {
		if err := json.Unmarshal(bytes.NewBufferString(body).Bytes(), &status); err != nil {
			return "", err
		}

		if status.responseCode == "00" {
			return status.status, nil
		}

		return "", fmt.Errorf("Could not confirm invoice status")
	}
}

// NewCheckoutInvoice - create a new checkout instance
//
// Example.
//     checkout := mpower.NewCheckoutInvoice(newSetup, newStore)
func NewCheckoutInvoice(setup *Setup, store *Store) *CheckoutInvoice {
	checkoutInvoiceIns := &CheckoutInvoice{Invoice: Invoice{Setup: setup, Store: *store}}
	checkoutInvoiceIns.baseURL = checkoutInvoiceIns.Invoice.Setup.BaseURL + "/checkout-invoice"
	return checkoutInvoiceIns
}
