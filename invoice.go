package mpowergo

import (
	"strconv"
)

type item struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float32 `json:"unit_price"`
	TotalPrice  float32 `json:"total_price"`
	Description string  `json:"description"`
}

type tax struct {
	Name   string  `json:"name"`
	Amount float32 `json:"amount"`
}

type invoice struct {
	itemsLen    int               `json:"-"`
	taxesLen    int               `json:"-"`
	Items       map[string]item   `json:"items"`
	Taxes       map[string]tax    `json:"taxes"`
	TotalAmount float32           `json:"total_amount"`
	Description string            `json:"description"`
	Actions     map[string]string `json:"actions"`
}

type Invoice struct {
	Setup      *Setup                 `json:"-"`
	Store      Store                  `json:"store"`
	InvoiceIn  invoice                `json:"invoice"`
	CustomData map[string]interface{} `json:"custom_data"`
}

func (i *Invoice) AddItem(name string, quantity int, unitPrice float32, totalPrice float32, desc string) {
	i.InvoiceIn.Items["item_"+strconv.Itoa(i.InvoiceIn.itemsLen)] = item{name, quantity, unitPrice, totalPrice, desc}
	i.InvoiceIn.itemsLen += 1
}

func (i *Invoice) AddTax(name string, amount float32) {
	i.InvoiceIn.Taxes["tax_"+strconv.Itoa(i.InvoiceIn.taxesLen)] = tax{name, amount}
	i.InvoiceIn.taxesLen += 1
}

func (i *Invoice) SetDescription(desc string) {
	if desc == "" {
		panic("provide the description argument")
	}

	i.InvoiceIn.Description = desc
}

func (i *Invoice) SetTotalAmount(amt float32) {
	if amt == 0 {
		panic("provide the totalAmount argument")
	}

	i.InvoiceIn.TotalAmount = amt
}

func (i *Invoice) SetCustomData(key string, val interface{}) {
	i.CustomData[key] = val
}
