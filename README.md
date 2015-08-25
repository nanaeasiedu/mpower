## MPOWER - GOLANG LIBRARY FOR MPOWER API

[![Build Status](https://secure.travis-ci.org/ngenerio/mpower.png?branch=master)](https://travis-ci.org/ngenerio/mpower)

[![GoDoc](https://godoc.org/github.com/ngenerio/mpower?status.svg)](https://godoc.org/github.com/ngenerio/mpower)

This a go implementation library that interfaces with the mpower http api.

Built on the MPower Payments [`HTTP API (beta)`](http://mpowerpayments.com/developers/http).

### Installation

```bash
$ go get github.com/ngenerio/mpower
```

### Documentation

Create a new store instance to use in the checkout or onsite invoice

```go
mpowerStore := NewStore("Awesome Store")
```

Create a new setup instance to use in the checkout or onsite invoice

```go
// Get your keys from MPower Integration Setup
mpowerSetup := NewSetup(MASTER_KEY, PRIVATE_KEY , PUBLIC_KEY, TOKEN)
```

#### Checkout and Onsite Invoice

To use the checkout invoice, you need create an mpower instance

```go
mpower := mpower.NewMPower(seup, store, "test")
checkout := mpower.NewCheckoutInvoice(mpower)
onsite := mpower.NewOnsiteInvoice(mpower)
```

Add an item to the invoice

```go
checkout.AddItem("Yam Phone", 1, 50.00, 50.00, "Hello World")
```

Add tax information to the invoice to be displayed on the cutomer's receipt

```go
checkout.AddTax("VAT", 30.00)
```

Set custom data on the invoice

```go
checkout.SetCustomData("bonus", yeah)
```

Set some description on the invoice

```go
checkout.SetDescription("Hello World")
```

Set the total amount on the invoice

```go
checkout.SetTotalAmount(80.00)
```

#### Creating an invoice

This sample code shows how to create an mpower invoice after adding some items to your `checkout` or `onsite` invoice

```go
//`response` is of type [`napping.Response`](http://godoc.org/github.com/jmcvetta/napping#Response)
responseBody, response, err := checkout.Create()

if err != nil {
    // handle the error
}
// where `responseBody.Token` is the token of the created invoice
fmt.Println(responseBody.Token)
```

#### Confirming the status of an invoice

```go
// `TOKEN` is the token of the invoice created
responseBody, response, err := checkout.Confirm(TOKEN)

if err != nil {
    // handle the error
}

// `response.Status` could either be `pending`, `cancelled` or `completed`
fmt.Println(responseBody.Status)
```

#### Charging the mpower customer with onsite payment request

```go
// `TOKEN` is the onsite token of the invoice created and the `CUSTOMER_TOKEN` is from the customer
responseBody, response, err := onsite.Charge(TOKEN, CUSTOMER_TOKEN)
```

#### Direct Mobile

Docs coming up soon

#### Direct Pay

Docs coming up soon

For more docs, read up:
[Mpower docs](https://godoc.org/github.com/ngenerio/mpower)
