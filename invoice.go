package mpower

import (
	"fmt"
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
	Taxes       map[string]tax    `json:"taxes,omitempty"`
	TotalAmount float32           `json:"total_amount"`
	Description string            `json:"description"`
	Actions     map[string]string `json:"actions,omitempty"`
}

type Invoice struct {
	Setup      *Setup                 `json:"-"`
	Store      Store                  `json:"store"`
	InvoiceIn  invoice                `json:"invoice"`
	CustomData map[string]interface{} `json:"custom_data,omitempty"`
}

func (i *Invoice) AddItem(name string, quantity int, unitPrice float32, totalPrice float32, desc string) {
	// check golang issue#3117 https://code.google.com/p/go/issues/detail?id=3117
	if i.InvoiceIn.itemsLen == 0 {
		i.InvoiceIn.Items = make(map[string]item)
	}
	itemName := "item_" + fmt.Sprintf("%d", i.InvoiceIn.itemsLen)
	i.InvoiceIn.Items[itemName] = item{}
	tempItem := i.InvoiceIn.Items[itemName]
	tempItem.Name = name
	tempItem.Quantity = quantity
	tempItem.UnitPrice = unitPrice
	tempItem.TotalPrice = totalPrice
	tempItem.Description = desc
	i.InvoiceIn.Items[itemName] = tempItem
	i.InvoiceIn.itemsLen += 1
}

func (i *Invoice) AddTax(name string, amount float32) {
	if i.InvoiceIn.taxesLen == 0 {
		i.InvoiceIn.Taxes = make(map[string]tax)
	}
	taxName := "tax_" + fmt.Sprintf("%d", i.InvoiceIn.taxesLen)
	i.InvoiceIn.Taxes[taxName] = tax{}
	tempTax := i.InvoiceIn.Taxes[taxName]
	tempTax.Name = name
	tempTax.Amount = amount
	i.InvoiceIn.Taxes[taxName] = tempTax
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
