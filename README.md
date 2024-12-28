# seerbit-go - SeerBit Payment API Library for GOLANG

## Features

The Library supports all APIs under the following services:

* Payments via API (mobile money, cards, account, etc.)
* Recurring Payments
* Transaction Status

## Getting Started

A full getting started guide for integrating SeerBit can be found at [getting started docs](https://doc.seerbit.com).

## Documentation

The documentation, installation guide, detailed description of the SeerBit API and all of its features is [available on the documentation website](https://doc.seerbit.com/api/library)

## Requirements

* GO 1.21.9

## Installation

Run this command on the terminal:

```code
go get https://github.com/hackdaemon2/seerbit-go
```

## Contributing

You can contribute to this repository so that anyone can benefit from it:

* Improved features
* Resolved bug fixes and issues

## Demos  

You can also check the [demos](https://github.com/hackdaemon2/seerbit-go/tree/master/demos)

## Using the Library

### Initiate Account Option

Instantiate a client and set the parameters.

```go
  package main

  import (
    "encoding/json"
    "fmt"
    "log"

    "github.com/hackdaemon2/seerbit-go/pkg/account"
    "github.com/hackdaemon2/seerbit-go/pkg/client"
    "github.com/hackdaemon2/seerbit-go/pkg/model"
  )

  func main() {
    processSeerBitPayment()
  }

  func processSeerBitPayment() {
    seerBitClient, err := client.NewSeerBitClient("your_public_key", "your_private_key")
    if err != nil {
      log.Fatalf("Failed to initialize SeerBit client: %v", err)
    }

    checkout := checkout.NewCheckout(seerBitClient)
    cardPayload := model.CheckoutPayload{
      PublicKey:          seerBitClient.PublicKey,
      Amount:             "100.00",
      Currency:           "NGN",
      Country:            "NG",
      PaymentReference:   "paymentRef",
      Email:              "johndoe@gmail.com",
      ProductId:          "Foods",
      ProductDescription: "Checkout Payment Transaction",
    }

    response, err := checkout.Pay(cardPayload)
    if err != nil {
      log.Fatalf("Error making payment: %v", err)
    }

    switch resp := response.(type) {
    case model.PaymentResponse:
      if resp.Data.Code == constant.SEERBIT_PENDING_CODE || resp.Data.Code == constant.SEERBIT_SUCCESS_CODE {
        printJson(resp)
        log.Printf("redirect link => %s", resp.Data.Payment.RedirectLink)
      } else {
        log.Printf("Payment failed: %v", resp)
      }
    case model.ErrorResponse:
      log.Printf("Payment failed: %v", resp)
    default:
      log.Printf("Unexpected response type: %T", resp)
    }
  }

  func printJson(response any) {
    jsonResponse, err := json.Marshal(response)
    if err != nil {
      log.Printf("Error marshalling JSON: %v", err)
      return
    }
    fmt.Println(string(jsonResponse))
  }

```

Find more examples [here](https://github.com/hackdaemon2/seerbit-go/tree/master/demos).

## Licence

The MIT License (MIT). For more information, see the LICENSE file.
