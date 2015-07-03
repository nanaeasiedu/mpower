package mpower

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
)

type CheckoutInvoice struct {
	Invoice
	baseUrl      string `json:"-"`
	ResponseCode string `json:"-"`
	ResponseText string `json:"-"`
	Description  string `json:"-"`
	Token        string `json:"-"`
	Status       string `json:"-"`
}

type responseJsonCheckout struct {
	ResponseCode string `json:"response_code"`
	ResponseText string `json:"response_text"`
	Description  string `json:"description"`
	Token        string `json:"token"`
}

type stat struct {
	status       string `json:"status"`
	responseCode string `json:"response_code"`
}

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

func CreateCheckoutInvoice(setup *Setup, store *Store) *CheckoutInvoice {
	checkoutInvoiceIns := &CheckoutInvoice{Invoice: Invoice{Setup: setup, Store: *store}}
	checkoutInvoiceIns.baseUrl = checkoutInvoiceIns.Invoice.Setup.BASE_URL + "/checkout-invoice"
	return checkoutInvoiceIns
}
