package mpower

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
)

// DirectPay - the direct pay object as defined by mpower
type DirectPay struct {
	baseUrl       string
	Setup         *Setup
	Status        string
	ResponseCode  string
	ResponseText  string
	Description   string
	TransactionId string
}

// payDta - `struct` to send the data as json to mpower
type payData struct {
	Alias  string `json:"account_alias"`
	Amount int    `json:"amount"`
}

// directResponse - the response from mpower is serialiazed into this form
type directPayReponse struct {
	RespnseCode   string `json:"response_code"`
	ResponseText  string `json:"response_text"`
	Description   string `json:"description"`
	TransactionId string `json:"transaction_id"`
}

// CreditAccount - credits the account of an mpower customer
//
// Example.
//    if ok, err := directPayInStance.CreditAccount("me", 500); ok {
//    everything was ok
//    } else {
//     There's trouble in hell
//    }
func (d *DirectPay) CreditAccount(account string, amount int) (bool, error) {
	dataToSend := payData{account, amount}
	req := gorequest.New()

	req.Post(d.baseUrl + "/credit-account")
	var dataToRecv directPayReponse

	for key, val := range d.Setup.Headers {
		req.Set(key, val)
	}

	if content, err := json.Marshal(dataToSend); err != nil {
		return false, err
	} else {
		req.Send(bytes.NewBuffer(content).String())
	}

	if resp, body, err := req.End(); err != nil {
		d.Status = resp.Status
		return false, fmt.Errorf("%v", err)
	} else {
		if err := json.Unmarshal(bytes.NewBufferString(body).Bytes(), &dataToRecv); err != nil {
			return false, err
		}

		if dataToRecv.RespnseCode == "00" {
			d.ResponseCode = dataToRecv.RespnseCode
			d.ResponseText = dataToRecv.ResponseText
			d.Description = dataToRecv.Description
			d.TransactionId = dataToRecv.TransactionId

			return true, nil
		}

		d.Status = resp.Status
		d.ResponseText = body
		return false, fmt.Errorf("Failed to to credit acoount %s with %d", account, amount)
	}
}

// NewDirectPay - creates a DirectPay instance
func NewDirectPay(setup *Setup) *DirectPay {
	directIns := &DirectPay{Setup: setup}
	directIns.baseUrl = directIns.Setup.BASE_URL + "/direct-pay"

	return directIns
}
