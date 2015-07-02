package mpowergo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
)

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
		req.Send(content)
	}

	if resp, body, err := req.End(); err != nil {
		on.Status = resp.Status
		return false, fmt.Errorf("%v", err)
	} else {
		if err := json.Unmarshal(bytes.NewBufferString(body).Bytes(), &respJson); err != nil {
			panic("Error decoding json")
		}

		if respJson.ResponseCode == "00" {
			on.ResponseText = respJson.ResponseText
			on.ResponseCode = respJson.ResponseCode
			on.Description = respJson.Description
			on.Token = respJson.Token
			on.InvoiceToken = respJson.InvoiceToken

			return true, nil
		}

		return false, fmt.Errorf("Failed to create invoice with error : %s", body)
	}
}

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
		if _, body, err := req.Send(dataByte).End(); err != nil {
			return false, fmt.Errorf("%v", err)
		} else {
			if err := json.Unmarshal(bytes.NewBufferString(body).Bytes(), &respData); err != nil {
				panic("Error decoding json")
			}

			if respData.ResponseCode == "00" {
				on.ResponseText = respData.ResponseText
				on.ResponseCode = respData.ResponseCode
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

func CreateOnsiteInvoice(setup Setup, store Store) *OnsiteInvoice {
	onsiteInvoiceIns := &OnsiteInvoice{Invoice: Invoice{Setup: &setup, Store: store}}
	onsiteInvoiceIns.baseUrl = onsiteInvoiceIns.Invoice.Setup.BASE_URL + "/opr"
	return onsiteInvoiceIns
}
