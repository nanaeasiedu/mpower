## MPOWERGO

[![Build Status](https://secure.travis-ci.org/ngenerio/mpowergo.png?branch=master)](https://travis-ci.org/ngenerio/mpowergo)

This a go implementation library that interfaces with the mpower http api.

Built on the MPower Payments [`HTTP API (beta)`](http://mpowerpayments.com/developers/http).

### Installation

```bash
$ go get github.com/ngenerio/mpowergo
```

### Documentation

Create a new store instance to use in the checkout or onsite invoice

```go
newSetup := mpower.NewStore(map[string]string{
    "name":          "Awesome Store",
    "tagline":       "Easy shopping",
    "phoneNumber":   "0272271893",
    "postalAddress": "P.0. Box MP555, Accra",
    "logoURL":       "http://www.awesomestore.com.gh/logo.png",
})
```

Create a new setup instance to use in the checkout or onsite invoice

```go
newSetup := mpower.NewSetup(map[string]string{
    "masterKey":  "55647970-22e1-4e7e-8fb4-56eca2b3b006",
    "privateKey": "test_private_B8EiE1AGWpb4tVMzVTyFDu9rYoc",
    "publicKey":  "test_public_B1wo2UVmxUrvwzZuPqpLrWqlA74",
    "token":      "a6d96e2586c8bbae7c28",
    "mode":       "test",
})
```

#### Checkout and Onsite Invoice 

To use the checkout invoice, you need your store and setup info above

```go
checkout := mpower.NewCheckoutInvoice(newSetup, newStore)
onsite := mpower.NewOnsiteInvoice(newSetup, newStore)
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


To create an invoice on mpower, call the **`Create`** method on the `checkout`

This is for the checkout invoice

```go
if ok, err := checkout.Create(); ok {
  //do something with the response info on the checkout instance
  fmt.Printf("%s %s %s %s\n\n", checkout.ResponseCode, checkout.ResponseText, checkout.Description, checkout.Token)
} else {
  //there was an error
}
```

For onsite invoice

```go
if ok, err := checkout.Create("me"); ok {
  //do something with the response info on the checkout instance
  fmt.Printf("%s %s %s %s\n\n", checkout.ResponseCode, checkout.ResponseText, checkout.Description, checkout.Token)
} else {
  //there was an error
}
```

Get the invoice url of the recently created invoice with **`GetInvoiceUrl`**

```go
str := checkout.GetInvoiceUrl()
```

For the onsite invoice, you have to charge the customer using the confirm token from **`Create`** and `token` from the user

```go
if str, err := checkout.Confirm(token); err != nil {
    // handle error
} else if str == "completed" {
    // do something
}
```
