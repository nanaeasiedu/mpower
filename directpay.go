package mpower

import (
	"github.com/jmcvetta/napping"
)

// DirectPay - the direct pay object as defined by mpower
type DirectPay struct {
	baseURL string
	mpower  *MPower
}

// DirectPayRequest - `struct` to send the data as json to mpower
type DirectPayRequest struct {
	Alias  string `json:"account_alias,omitempty"`
	Amount int    `json:"amount,omitempty"`
}

// DirectPayResponse - the response from mpower is serialiazed into this form
type DirectPayResponse struct {
	Response
	TransactionID string `json:"transaction_id,omitempty"`
}

// CreditAccount - credits the account of an mpower customer
func (d *DirectPay) CreditAccount(account string, amount int) (*DirectPayResponse, *napping.Response, error) {
	payload := &DirectPayRequest{account, amount}
	responseBody := &DirectPayResponse{}

	response, err := d.mpower.NewRequest("POST", d.baseURL+"/credit-account", payload, responseBody, nil)

	if err != nil || responseBody.ResponseCode != "00" {
		return nil, response, err
	}

	return responseBody, response, nil
}

// NewDirectPay - creates a DirectPay instance
func NewDirectPay(mp *MPower) *DirectPay {
	return &DirectPay{
		mpower:  mp,
		baseURL: mp.baseURL + "/direct-pay",
	}
}
