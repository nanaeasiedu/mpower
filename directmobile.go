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

// DirectMobileResponse is the response from a direct mobile charge request
type DirectMobileResponse struct {
	Response
	Token               string `json:"token"`
	TransactionID       string `json:"transaction_id"`
	MobileInvoiceNumber string `json:"mobile_invoice_no"`
}

// DirectMobileStatusResponse is the status of a direct mobile transaction
type DirectMobileStatusResponse struct {
	Response
	TXStatus            string `json:"tx_status"`
	TransactionID       string `json:"transaction_id"`
	MobileInvoiceNumber string `json:"mobile_invoice_no"`
	CancelReason        string `json:"cancel_reason"`
}

func (d *DirectMobile) directMobileLiveHeader() *http.Header {
	header := new(http.Header)
	header.Add("MP-Master-Key", d.mpower.setup.MasterKey)
	header.Add("MP-Private-Key", d.mpower.setup.PrivateKey)
	header.Add("MP-Token", d.mpower.setup.Token)

	return header
}

// Charge charges customers' mobile money money wallets directly on your site or application
//
// Example.
//		resp, err := directMobileInstance.Charge("Eugene", "ngene84@gmail.com", "0272271893", "Awesome Shopping", "MTN", "20")
func (d *DirectMobile) Charge(name, email, phone, merchant, wallet, amount string) (*DirectMobileResponse, *napping.Response, error) {
	if d.mpower.mode == "test" {
		panic("Cannot make a direct mobile requset in `test` mode")
	}
	payload := &DirectMobileRequest{name, email, phone, merchant, wallet, amount}
	responseBody := &DirectMobileResponse{}

	response, err := d.mpower.NewRequest("POST", d.baseURL+"/charge", payload, responseBody, d.directMobileLiveHeader())

	if err != nil || responseBody.ResponseCode != "00" {
		return nil, response, err
	}

	return responseBody, response, nil
}

// Status checks the status of a direct mobile transaction
func (d *DirectMobile) Status(token string) (*DirectMobileStatusResponse, *napping.Response, error) {
	if d.mpower.mode == "test" {
		panic("Cannot make a direct mobile requset in `test` mode")
	}
	payload := &struct {
		token string
	}{token}
	responseBody := &DirectMobileStatusResponse{}

	response, err := d.mpower.NewRequest("POST", d.baseURL+"/status", payload, responseBody, d.directMobileLiveHeader())

	if err != nil || responseBody.ResponseCode != "00" {
		return nil, response, err
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
