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
	baseUrl      string `json:"-"`
	ResponseCode string `json:"-"`
	ResponseText string `json:"-"`
	Description  string `json:"-"`
	Token        string `json:"-"`
	Status       string `json:"-"`
}

// The response data as specified by the mpower
// It retrieves the response json data and stores it on the checkout invoice object
type responseJsonCheckout struct {
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
	var respJson responseJsonCheckout
	req := gorequest.New()

	req.Post(c.baseUrl + "/create")

	for key, val := range c.Setup.GetHeaders() {
		req.Set(key, val)
	}

	if content, err := json.Marshal(c.Invoice); err != nil {
		panic("Error encoding json")
	} else {
		req.Send(bytes.NewBuffer(content).String())
	}

	if resp, body, err := req.End(); err != nil {
		fmt.Errorf("%v", err)
		c.Status = resp.Status
		return false, fmt.Errorf("%v", err)
	} else {
		if err := json.Unmarshal(bytes.NewBufferString(body).Bytes(), &respJson); err != nil {
			panic("Error decoding json")
		} else if respJson.ResponseCode == "00" {
			c.ResponseText = respJson.ResponseText
			c.ResponseCode = respJson.ResponseCode
			c.Description = respJson.Description
			c.Token = respJson.Token

			return true, nil
		}

		return false, fmt.Errorf("Failed to create invoice with error : %s", body)
	}
}

// GetInvoiceUrl - get the invoice's url from the response
//
// Example.
//    str := checkout.GetInvoiceUrl()
func (c *CheckoutInvoice) GetInvoiceUrl() string {
	if c.Token == "" {
		panic("Token currently not available")
	}

	return c.ResponseText
}

func (c *CheckoutInvoice) Confirm(token string) (string, error) {
	var status stat
	req := gorequest.New()

	req.Get(c.baseUrl + "/confirm/" + token)
	for key, val := range c.Setup.GetHeaders() {
		req.Set(key, val)
	}

	if _, body, err := req.End(); err != nil {
		return "", fmt.Errorf("%v", err)
	} else {
		if err := json.Unmarshal(bytes.NewBufferString(body).Bytes(), &status); err != nil {
			return "", err
		}

		if status.responseCode == "00" {
			return status.status, nil
		} else {
			return "", fmt.Errorf("Could not confirm invoice status")
		}
	}
}

// NewCheckoutInvoice - create a new checkout instance
//
// Example.
//     checkout := mpower.NewCheckoutInvoice(newSetup, newStore)
func NewCheckoutInvoice(setup *Setup, store *Store) *CheckoutInvoice {
	checkoutInvoiceIns := &CheckoutInvoice{Invoice: Invoice{Setup: setup, Store: *store}}
	checkoutInvoiceIns.baseUrl = checkoutInvoiceIns.Invoice.Setup.BASE_URL + "/checkout-invoice"
	return checkoutInvoiceIns
}
