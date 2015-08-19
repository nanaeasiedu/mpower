## MPOWERGO

[![Build Status](https://secure.travis-ci.org/ngenerio/mpowergo.png?branch=master)](https://travis-ci.org/ngenerio/mpowergo)

[![GoDoc](https://godoc.org/github.com/ngenerio/mpowergo?status.svg)](https://godoc.org/github.com/ngenerio/mpowergo)

This a go implementation library that interfaces with the mpower http api.

Built on the MPower Payments [`HTTP API (beta)`](http://mpowerpayments.com/developers/http).

### Installation

```bash
$ go get github.com/ngenerio/mpowergo
```

### Documentation

Create a new store instance to use in the checkout or onsite invoice

```go
mpowerStore := NewStore("Awesome Store", "Easy shopping", "0272271893", "P.0. Box MP555, Accra", "http://www.awesomestore.com.gh/logo.png")
```

Create a new setup instance to use in the checkout or onsite invoice

```go
mpowerSetup := NewSetup("43434-54545-45454-545432", "test_private_auhidaudvbirbyyrieoib", "test_public_iopjasdioppdadipjoasd", "ioapdojdifouw8h")
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
For more docs, read up:
[Mpowergo docs](https://godoc.org/github.com/ngenerio/mpowergo)
