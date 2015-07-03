package mpower

import (
	"fmt"
)

// Item definition as specified by mpower docs
// It holds the data of an item
type item struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float32 `json:"unit_price"`
	TotalPrice  float32 `json:"total_price"`
	Description string  `json:"description"`
}

// Tax definition as specified by mpower docs
// It holds the tax data
type tax struct {
	Name   string  `json:"name"`
	Amount float32 `json:"amount"`
}

// Invoice definition as specified by mpower docs
// It holds all the data related to the invoice
type invoice struct {
	itemsLen    int               `json:"-"`
	taxesLen    int               `json:"-"`
	Items       map[string]item   `json:"items"`
	Taxes       map[string]tax    `json:"taxes,omitempty"`
	TotalAmount float32           `json:"total_amount"`
	Description string            `json:"description"`
	Actions     map[string]string `json:"actions,omitempty"`
}

// The invoice definition
// It specifies the required field keys and values we will be sending over to mpower
// This is supposed to be an embedded struct in the Onsite Invoice and Checkout Invoice
type Invoice struct {
	Setup      *Setup                 `json:"-"`
	Store      Store                  `json:"store"`
	InvoiceIn  invoice                `json:"invoice"`
	CustomData map[string]interface{} `json:"custom_data,omitempty"`
}

// AddItem add an `item - struct` to the items in the invoice
//
// Example.
//    checkout := mpower.NewCheckoutInvoice(newSetup, newStore)
//    checkout.AddItem("Yam Phone", 1, 50.00, 50.00, "Hello World")
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

// AddItem add an `tax - struct` to the taxes in the invoice
//
// Example.
//    checkout := mpower.NewCheckoutInvoice(newSetup, newStore)
//    checkout.AddTax("VAT", 30.00)
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

// Sets the description for the invoice
//
// Example.
//    checkout := mpower.NewCheckoutInvoice(newSetup, newStore)
//    checkout.SetDescription("Hello World")
func (i *Invoice) SetDescription(desc string) {
	if desc == "" {
		panic("provide the description argument")
	}

	i.InvoiceIn.Description = desc
}

// Sets the total amount on the invoice
//
// Example.
//    checkout := mpower.NewCheckoutInvoice(newSetup, newStore)
//    checkout.SetTotalAmount(80.00)
func (i *Invoice) SetTotalAmount(amt float32) {
	if amt == 0 {
		panic("provide the totalAmount argument")
	}

	i.InvoiceIn.TotalAmount = amt
}

// Sets the total amount on the invoice
//
// Example.
//    checkout := mpower.NewCheckoutInvoice(newSetup, newStore)
//    checkout.SetCustomData("bonus", yeah)
func (i *Invoice) SetCustomData(key string, val interface{}) {
	i.CustomData[key] = val
}
