package mpower

import (
	"fmt"
	"sync"
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
	ItemsArr    []item            `json:"-"`
	TaxesArr    []tax             `json:"-"`
	Items       map[string]item   `json:"items"`
	Taxes       map[string]tax    `json:"taxes,omitempty"`
	TotalAmount float32           `json:"total_amount"`
	Description string            `json:"description"`
	Actions     map[string]string `json:"actions,omitempty"`
}

// Invoice definition
// It specifies the required field keys and values we will be sending over to mpower
// This is supposed to be an embedded struct in the Onsite Invoice and Checkout Invoice
type Invoice struct {
	sync.RWMutex
	Setup       *Setup                 `json:"-"`
	Store       Store                  `json:"store"`
	InvoiceData invoice                `json:"invoice"`
	CustomData  map[string]interface{} `json:"custom_data,omitempty"`
}

// AddItem add an `item - struct` to the items in the invoice
//
// Example.
//    checkout := mpower.NewCheckoutInvoice(newSetup, newStore)
//    checkout.AddItem("Yam Phone", 1, 50.00, 50.00, "Hello World")
func (i *Invoice) AddItem(name string, quantity int, unitPrice float32, totalPrice float32, desc string) error {
	for _, value := range i.InvoiceData.ItemsArr {
		if value.Name == name {
			return fmt.Errorf("Invoice item with name %s already exists", name)
		}
	}
	tempItem := item{}
	tempItem.Name = name
	tempItem.Quantity = quantity
	tempItem.UnitPrice = unitPrice
	tempItem.TotalPrice = totalPrice
	tempItem.Description = desc

	i.InvoiceData.ItemsArr = append(i.InvoiceData.ItemsArr, tempItem)
	return nil
}

// RemoveItem removes the item with name of `name`
//
// Example.
//     checkout.RemoveItem()
func (i *Invoice) RemoveItem(name string) {
	for ix, value := range i.InvoiceData.ItemsArr {
		if value.Name == name {
			i.InvoiceData.ItemsArr = append(i.InvoiceData.ItemsArr[:ix], i.InvoiceData.ItemsArr[ix+1:]...)
			break
		}
	}
}

// ClearAllItems clears all the items in the invoice
//
// Example.
//     checkout.ClearAllItems()
func (i *Invoice) ClearAllItems() {
	i.InvoiceData.ItemsArr = nil
	i.InvoiceData.Items = make(map[string]item)
}

// AddTax add an `tax - struct` to the taxes in the invoice
//
// Example.
//    checkout := mpower.NewCheckoutInvoice(newSetup, newStore)
//    checkout.AddTax("VAT", 30.00)
func (i *Invoice) AddTax(name string, amount float32) error {
	for _, value := range i.InvoiceData.TaxesArr {
		if value.Name == name {
			return fmt.Errorf("Tax with %s already exists", name)
		}
	}
	tempTax := tax{}
	tempTax.Name = name
	tempTax.Amount = amount

	i.InvoiceData.TaxesArr = append(i.InvoiceData.TaxesArr, tempTax)
	return nil
}

// RemoveTax removes the tax with name of `name`
//
// Example.
//     checkout.RemoveTax()
func (i *Invoice) RemoveTax(name string) {
	for ix, value := range i.InvoiceData.TaxesArr {
		if value.Name == name {
			i.InvoiceData.TaxesArr = append(i.InvoiceData.TaxesArr[:ix], i.InvoiceData.TaxesArr[ix+1:]...)
			break
		}
	}
}

// ClearAllTaxes clears all the taxes in the invoice
//
// Example.
//     checkout.ClearAllTaxes()
func (i *Invoice) ClearAllTaxes() {
	i.InvoiceData.TaxesArr = nil
	i.InvoiceData.Taxes = make(map[string]tax)
}

// Clear clears all the items in the invoice
//
// Example.
//     checkout.Clear()
func (i *Invoice) Clear() {
	i.ClearAllItems()
	i.ClearAllTaxes()
}

// SetDescription the description for the invoice
//
// Example.
//    checkout := mpower.NewCheckoutInvoice(newSetup, newStore)
//    checkout.SetDescription("Hello World")
func (i *Invoice) SetDescription(desc string) {
	if desc == "" {
		panic("provide the description argument")
	}

	i.InvoiceData.Description = desc
}

// SetTotalAmount the total amount on the invoice
//
// Example.
//    checkout := mpower.NewCheckoutInvoice(newSetup, newStore)
//    checkout.SetTotalAmount(80.00)
func (i *Invoice) SetTotalAmount(amt float32) {
	if amt == 0 {
		panic("provide the totalAmount argument")
	}

	i.InvoiceData.TotalAmount = amt
}

// SetCustomData the total amount on the invoice
//
// Example.
//    checkout := mpower.NewCheckoutInvoice(newSetup, newStore)
//    checkout.SetCustomData("bonus", yeah)
func (i *Invoice) SetCustomData(key string, val interface{}) {
	if i.CustomData == nil {
		i.CustomData = make(map[string]interface{})
	}
	i.Lock()
	i.CustomData[key] = val
	i.Unlock()
}

// PrepareForRequest prepares the invoice for request
// This is called before the request to mpower invoice is made to set the items and taxes into a json format
func (i *Invoice) PrepareForRequest() {
	i.InvoiceData.Items = make(map[string]item)
	i.InvoiceData.Taxes = make(map[string]tax)

	// Check the section on `concurrrency` http://blog.golang.org/go-maps-in-action
	// http://golang.org/doc/faq#atomic_maps
	i.Lock()

	for ix, value := range i.InvoiceData.ItemsArr {
		itemName := fmt.Sprintf("item_%d", ix)
		i.InvoiceData.Items[itemName] = item{}
		tempItem := i.InvoiceData.Items[itemName]
		tempItem.Name = value.Name
		tempItem.Quantity = value.Quantity
		tempItem.UnitPrice = value.UnitPrice
		tempItem.TotalPrice = value.TotalPrice
		tempItem.Description = value.Description
		i.InvoiceData.Items[itemName] = tempItem
	}

	for ix, value := range i.InvoiceData.TaxesArr {
		taxName := fmt.Sprintf("tax_%d", ix)
		i.InvoiceData.Taxes[taxName] = tax{}
		tempTax := i.InvoiceData.Taxes[taxName]
		tempTax.Name = value.Name
		tempTax.Amount = value.Amount
		i.InvoiceData.Taxes[taxName] = tempTax
	}

	i.Unlock()
}
