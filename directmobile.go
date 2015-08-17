package mpower

import (
	"github.com/jmcvetta/napping"
	"net/http"
)

// DirectMobile is used to handle api requests to mpower direct mobile payments
type DirectMobile struct {
	mpower  *MPower
	baseURL string
}

// DirectMobileRequest The request json to be sent over during a `charge` api request
type DirectMobileRequest struct {
	CustomerName   string `json:"customer_name"`
	CustomerEmail  string `json:"customer_email"`
	CustomerPhone  string `json:"customer_phone"`
	MerchantName   string `json:"merchant_name"`
	WalletProvider string `json:"wallet_provider"`
	Amount         string `json:"amount"`
}

type DirectMobileResponse struct {
	Response
	Token               string `json:"token"`
	TransactionID       string `json:"transaction_id"`
	MobileInvoiceNumber string `json:"mobile_invoice_no"`
}

type DirectMobileStatusResponse struct {
	Response
	TXStatus            string `json:"tx_status"`
	TransactionID       string `json:"transaction_id"`
	MobileInvoiceNumber string `json:"mobile_invoice_no"`
	CancelReason        string `json:"cancel_reason"`
}

// Charge charges customers' mobile money money wallets directly on your site or application
//
// Example.
//		resp, err := directMobileInstance.Charge("Eugene", "ngene84@gmail.com", "0272271893", "MTN", "20")
func (d *DirectMobile) Charge(name, email, phone, merchant, wallet, amount string) (*DirectMobileResponse, *napping.Response, error) {
	payload := &DirectMobileRequest{name, email, phone, merchant, wallet, amount}
	responseBody := &DirectMobileResponse{}

	header := new(http.Header)
	header.Add("MP-Master-Key", d.mpower.setup.MasterKey)
	header.Add("MP-Private-Key", d.mpower.setup.PrivateKey)
	header.Add("MP-Token", d.mpower.setup.Token)

	response, err := d.mpower.NewRequest("POST", d.baseURL+"/charge", payload, responseBody, header)

	if err != nil || responseBody.ResponseCode != "00" {
		return nil, response, err
	}

	return responseBody, response, nil
}

func (d *DirectMobile) Status(token string) (*DirectMobileStatusResponse, *napping.Response, error) {
	payload := &struct {
		token string
	}{token}
	responseBody := &DirectMobileStatusResponse{}

	header := new(http.Header)
	header.Add("MP-Master-Key", d.mpower.setup.MasterKey)
	header.Add("MP-Private-Key", d.mpower.setup.PrivateKey)
	header.Add("MP-Token", d.mpower.setup.Token)

	response, err := d.mpower.NewRequest("POST", d.baseURL+"/status", payload, responseBody, header)

	if err != nil || responseBody.ResponseCode != "00" {
		return nil, nil, err
	}

	return responseBody, response, nil

}

// NewDirectMobile creates a new DirectMobile instance
func NewDirectMobile(mp *MPower) *DirectMobile {
	return &DirectMobile{
		mpower:  mp,
		baseURL: mp.baseURL + "/direct-mobile",
	}
}
