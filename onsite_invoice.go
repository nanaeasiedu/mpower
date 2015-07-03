package mpower

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
)

// The onsite definition as defined by mpower documentation
// This struct holds all the data with respect to onsite request
type OnsiteInvoice struct {
	Invoice `json:"invoice_data"`
	OPRData struct {
		Alias string `json:"account_alias"`
	} `json:"opr_data"`
	baseUrl      string `json:"-"`
	ReceiptUrl   string `json:"-"`
	ResponseCode string `json:"-"`
	ResponseText string `json:"-"`
	Description  string `json:"-"`
	Token        string `json:"-"`
	InvoiceToken string `json:"-"`
	Status       string `json:"-"`
	Customer     struct {
		Name  string `json:"-"`
		Phone string `json:"-"`
		Email string `json:"-"`
	} `json:"-"`
}

type responseJsonOnsite struct {
	ResponseCode string `json:"response_code"`
	ResponseText string `json:"response_text"`
	Description  string `json:"description"`
	Token        string `json:"token"`
	InvoiceToken string `json:"invoice_token"`
}

type oprResponse struct {
	ResponseCode string `json:"response_code"`
	ResponseText string `json:"response_text"`
	Description  string `json:"description"`
	InvoiceData  struct {
		ReceiptUrl string `json:"receipt_url"`
		Status     string `json:"status"`
		Customer   struct {
			Name  string `json:"name"`
			Phone string `json:"phone"`
			Email string `json:"email"`
		} `json:"customer"`
	} `json:"invoice_data"`
}

type opr struct {
	token        string `json:"token"`
	confirmToken string `json:"confirm_token"`
}

// Create - creates a bew invoice on mpowers server
// Returns a `boolean` and `error`
// The boolean signifies whether the inovoice was created on not
// The response json object can be retrieved on the onsite invoice object created by the NewOnsiteInvoice
//
// Example.
//
//   ok, err := onsite.Create("hello")
//      if ok {
//         fmt.Printf("%s %s %s %s", onsite.ResponseCode, onsite.ResponseText, onsite.Description, onsite.Token)
//      } else {
//          fmt.Printf("%v", err)
//      }
func (on *OnsiteInvoice) Create(name string) (bool, error) {
	var respJson responseJsonOnsite
	req := gorequest.New()

	req.Post(on.baseUrl + "/create")

	for key, val := range on.Setup.GetHeaders() {
		req.Set(key, val)
	}

	on.OPRData.Alias = name
	if content, err := json.Marshal(on); err != nil {
		panic("Error encoding json")
	} else {
		req.Send(bytes.NewBuffer(content).String())
	}

	if resp, body, err := req.End(); err != nil {
		on.Status = resp.Status
		return false, fmt.Errorf("%v", err)
	} else {
		if err := json.Unmarshal(bytes.NewBufferString(body).Bytes(), &respJson); err != nil {
			panic("Error decoding json")
		}

		on.ResponseText = respJson.ResponseText
		on.ResponseCode = respJson.ResponseCode
		if respJson.ResponseCode == "00" {
			on.Description = respJson.Description
			on.Token = respJson.Token
			on.InvoiceToken = respJson.InvoiceToken

			return true, nil
		}

		return false, fmt.Errorf("Failed to create invoice with error : %s", body)
	}
}

// Charge - it charges the customer on mpower and returns a response json object which contains the receipt url with other information
// The `confirmToken` is from the customer
// Returns a `boolean` and `error`
// The boolean signifies whether the customer was chargeed or not
// The response json object can be retrieved on the onsite invoice object
//
// Example.
//    if ok, err := onsite.Charge(onsite.Token, "4346"); ok {
//      //doSomething
//    } else {
//
//    }
//
func (on *OnsiteInvoice) Charge(oprToken, confirmToken string) (bool, error) {
	var respData oprResponse
	data := opr{oprToken, confirmToken}
	req := gorequest.New()

	for key, val := range on.Setup.GetHeaders() {
		req.Set(key, val)
	}

	if dataByte, err := json.Marshal(data); err != nil {
		panic("Error encoding struct data")
	} else {
		if _, body, err := req.Send(bytes.NewBuffer(dataByte).String()).End(); err != nil {
			return false, fmt.Errorf("%v", err)
		} else {
			if err := json.Unmarshal(bytes.NewBufferString(body).Bytes(), &respData); err != nil {
				panic("Error decoding json")
			}

			on.ResponseText = respData.ResponseText
			on.ResponseCode = respData.ResponseCode

			if respData.ResponseCode == "00" {
				on.Description = respData.Description
				on.Status = respData.InvoiceData.Status
				on.ReceiptUrl = respData.InvoiceData.ReceiptUrl
				return true, nil
			} else {
				return false, fmt.Errorf("Failed to charge invoice. Check OPR or confirm token and try again.")
			}
		}
	}
}

// NewOnsiteInvoice create a new onsite invoice object
// It require a setup and store object
//
// Example.
//    onsite := mpower.NewOnsiteInvoice(newSetup, newStore)
func NewOnsiteInvoice(setup *Setup, store *Store) *OnsiteInvoice {
	onsiteInvoiceIns := &OnsiteInvoice{Invoice: Invoice{Setup: setup, Store: *store}}
	onsiteInvoiceIns.baseUrl = onsiteInvoiceIns.Invoice.Setup.BASE_URL + "/opr"
	return onsiteInvoiceIns
}
