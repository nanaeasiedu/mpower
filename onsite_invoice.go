package mpower

import (
	"github.com/jmcvetta/napping"
)

type OnsiteInvoice struct {
	Invoice
	baseURL string
	mpower  *MPower
}

// OnsiteInvoiceRequest
// The onsite definition as defined by mpower documentation
// This struct holds all the data with respect to onsite request
type OnsiteInvoiceRequest struct {
	Invoice                  `json:"invoice_data"`
	OnsitePaymentRequestData struct {
		Alias string `json:"account_alias"`
	} `json:"opr_data"`
}

type OnsiteInvoiceResponse struct {
	Response
	Token        string `json:"token"`
	InvoiceToken string `json:"invoice_token"`
}

type OnsitePaymentRequestCharge struct {
	Token        string `json:"token"`
	ConfirmToken string `json:"confirm_token"`
}

type OnsitePaymentRequestChargeResponse struct {
	InvoiceData struct {
		ReceiptURL string `json:"receipt_url"`
		Status     string `json:"status"`
		Invoice    struct {
			TotalAmount float32 `json:"total_amount"`
			Description string  `json:"description"`
		}
		Customer struct {
			Name  string `json:"name"`
			Phone string `json:"phone"`
			Email string `json:"email"`
		} `json:"customer"`
	} `json:"invoice_data"`
}

// Create - creates a new invoice on mpowers server
func (on *OnsiteInvoice) Create(name string) (*OnsiteInvoiceResponse, *napping.Response, error) {
	on.PrepareForRequest()

	requestBody := &OnsiteInvoiceRequest{}
	requestBody.Invoice = on.Invoice
	requestBody.OnsitePaymentRequestData.Alias = name

	responseBody := &OnsiteInvoiceResponse{}

	resp, err := on.mpower.NewRequest("POST", on.baseURL+"/create", requestBody, responseBody, nil)

	if err != nil {
		return nil, nil, err
	}

	return responseBody, resp, err
}

// Charge - it charges the customer on mpower and returns a response json object which contains the receipt url with other information
// The `confirmToken` is from the customer
func (on *OnsiteInvoice) Charge(onsitePaymentRequestToken, customerConfirmToken string) (*OnsitePaymentRequestChargeResponse, *napping.Response, error) {
	payload := &OnsitePaymentRequestCharge{onsitePaymentRequestToken, customerConfirmToken}
	responseBody := &OnsitePaymentRequestChargeResponse{}

	resp, err := on.mpower.NewRequest("POST", on.baseURL+"/charge", payload, responseBody, nil)

	if err != nil {
		return nil, nil, err
	}

	return responseBody, resp, err
}

// NewOnsiteInvoice create a new onsite invoice object
// It require a setup and store object
//
// Example.
//    onsite := mpower.NewOnsiteInvoice(newSetup, newStore)
func NewOnsiteInvoice(mp *MPower) *OnsiteInvoice {
	onsiteInvoice := &OnsiteInvoice{Invoice: Invoice{Setup: mp.setup, Store: *mp.store}}
	onsiteInvoice.baseURL = mp.baseURL + "/opr"
	return onsiteInvoice
}
